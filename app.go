package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router      *mux.Router
	ClientTheDP *mongo.Client
	Client34st  *mongo.Client
	ClientUTB   *mongo.Client
	TheDP       *mongo.Collection
	_34st       *mongo.Collection
	UTB         *mongo.Collection
}

func (a *App) Initialize(user, password string) {
	log.Default().Print("Connecting to MongoDB...")
	// database connection
	if user == "" || password == "" {
		log.Fatal("Missing required environment variables")
	}

	uri_thedp := fmt.Sprintf("mongodb+srv://%s:%s@dp.5aehsyo.mongodb.net/?retryWrites=true&w=majority", user, password)
	uri_34st := fmt.Sprintf("mongodb+srv://%s:%s@34st.rvmlxes.mongodb.net/?retryWrites=true&w=majority", user, password)
	uri_utb := fmt.Sprintf("mongodb+srv://%s:%s@utb.rjdlubs.mongodb.net/?retryWrites=true&w=majority", user, password)

	var client *mongo.Client
	var err error

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri_thedp))
	if err != nil {
		log.Fatal(err)
	}
	a.ClientTheDP = client
	a.TheDP = a.ClientTheDP.Database("Cluster").Collection("articles")

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri_34st))
	if err != nil {
		log.Fatal(err)
	}
	a.Client34st = client
	a._34st = a.Client34st.Database("Cluster").Collection("articles")

	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri_utb))
	if err != nil {
		log.Fatal(err)
	}
	a.ClientUTB = client
	a.UTB = a.ClientUTB.Database("Cluster").Collection("articles")

	log.Default().Println("Connected!")

	// routes
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/health", a.getHealth).Methods("GET")

	log.Default().Println("Routes initialized.")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

	defer func() {
		if err := a.ClientTheDP.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		if err := a.Client34st.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
		if err := a.ClientUTB.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
}

func (a *App) getHealth(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var result bson.M
	if err := a.ClientTheDP.Database("Cluster").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		Utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Error connection to TheDP database"})
		panic(err)
	}
	if err := a.Client34st.Database("Cluster").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		Utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Error connection to 34st database"})
		panic(err)
	}
	if err := a.ClientUTB.Database("Cluster").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		Utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Error connection to UTB database"})
		panic(err)
	}

	Utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Up and running!"})

	log.Default().Println("Health check successful.")
}
