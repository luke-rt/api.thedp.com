package main_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"thedp.com/api/api"
)

var a api.App

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	a = api.App{}
	a.Initialize(
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
	)

	code := m.Run()

	os.Exit(code)
}

func TestGetHealth(t *testing.T) {
	t.Logf("Testing GET /health")

	req, _ := http.NewRequest("GET", "/health", nil)
	res := executeRequest(t, req)

	assertOK(t, res)
	assertEqualJson(t, res, map[string]string{"message": "api.thedp.com: Up and running!"})
}

func TestGetRecent(t *testing.T) {
	t.Logf("Testing GET /dp/articles/recent/3")

	req, _ := http.NewRequest("GET", "/dp/articles/recent/3", nil)
	res := executeRequest(t, req)

	assertOK(t, res)

	var articles []api.Article
	err := json.Unmarshal(res.Body.Bytes(), &articles)

	if err != nil {
		t.Errorf("Expected response to be an array of articles. Got %s\n", res.Body.String())
	} else {
		t.Log("Expected response to be an array of articles. Success")
	}

	if len(articles) != 3 {
		t.Errorf("Expected 3 articles. Got %d\n", len(articles))
	} else {
		t.Log("Expected 3 articles. Success")
	}
}

func executeRequest(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()

	rec := httptest.NewRecorder()
	a.Router.ServeHTTP(rec, req)

	return rec
}

func assertOK(t *testing.T, res *httptest.ResponseRecorder) {
	if res.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.Code)
	} else {
		t.Logf("Expected response code %d. Success", http.StatusOK)
	}
}

func assertEqualJson(t *testing.T, res *httptest.ResponseRecorder, expected interface{}) {
	expectedJson, _ := json.Marshal(expected)
	if res.Body.String() != string(expectedJson) {
		t.Errorf("Expected response %s. Got %s\n", string(expectedJson), res.Body.String())
	} else {
		t.Logf("Expected response %s. Success", string(expectedJson))
	}
}
