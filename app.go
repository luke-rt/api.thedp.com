package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router      *mux.Router
	clients     map[string]*mongo.Client
	collections map[string]*mongo.Collection
	uris        map[string]string
}

func (a *App) Initialize(user, password string) {
	log.Default().Print("Connecting to MongoDB...")
	// database connection
	if user == "" || password == "" {
		log.Fatal("Missing required environment variables")
	}

	a.uris = make(map[string]string)
	a.clients = make(map[string]*mongo.Client)
	a.collections = make(map[string]*mongo.Collection)

	a.uris["dp"] = fmt.Sprintf("mongodb+srv://%s:%s@dp.5aehsyo.mongodb.net/?retryWrites=true&w=majority", user, password)
	a.uris["34st"] = fmt.Sprintf("mongodb+srv://%s:%s@34st.rvmlxes.mongodb.net/?retryWrites=true&w=majority", user, password)
	a.uris["utb"] = fmt.Sprintf("mongodb+srv://%s:%s@utb.rjdlubs.mongodb.net/?retryWrites=true&w=majority", user, password)

	for key, uri := range a.uris {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
		a.clients[key] = client
		a.collections[key] = client.Database("Cluster").Collection("articles")
	}

	log.Default().Println("Connected!")

	// routes
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/health", utils.Log(a.getHealth)).Methods("GET")
	a.Router.HandleFunc("/{db}/articles/recent/{count}", utils.Log(a.getRecent)).Methods("GET")

	log.Default().Println("Routes initialized.")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

	defer func() {
		for _, client := range a.clients {
			if err := client.Disconnect(context.TODO()); err != nil {
				log.Fatal(err)
			}
		}
	}()
}

func (a *App) getHealth(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var result bson.M
	for _, client := range a.clients {
		if err := client.Database("Cluster").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
			utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Error connection to TheDP database"})
			log.Fatal(err)
		}
	}

	utils.Json(res, http.StatusOK, map[string]string{"message": "api.thedp.com: Up and running!"})
}

func (a *App) getRecent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	db := vars["db"]
	count, err := strconv.Atoi(vars["count"])
	if err != nil {
		utils.Json(res, http.StatusBadRequest, map[string]string{"message": "api.thedp.com: Invalid count"})
		return
	}

	filter := bson.D{}
	opts := options.Find().SetLimit(int64(count))
	cursor, err := a.collections[db].Find(context.Background(), filter, opts)
	if err != nil {
		utils.Json(res, http.StatusInternalServerError, map[string]string{"message": "api.thedp.com: Error retrieving articles"})
		return
	}

	var results []Article
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	utils.Json(res, http.StatusOK, results)
}
