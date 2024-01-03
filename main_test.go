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

	assertSuccess(t, res)
	assertEqualJson(t, res, map[string]string{"message": "api.thedp.com: Up and running!"})
}

func TestLimit(t *testing.T) {
	t.Logf("Testing GET /dp/articles?limit=3")

	req, _ := http.NewRequest("GET", "/dp/articles?limit=3", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)

	articles := extractArticles(t, res)
	if len(articles) == 3 {
		t.Log("Expected 3 articles. Success")
	} else {
		t.Errorf("Expected 3 articles. Got %d\n", len(articles))
	}
}

func TestInvalidLimit(t *testing.T) {
	t.Logf("Testing GET /dp/articles?limit=nan")

	req, _ := http.NewRequest("GET", "/dp/articles?limit=nan", nil)
	res := executeRequest(t, req)

	assertStatus(t, res, http.StatusBadRequest)
	assertEqualJson(t, res, map[string]string{"message": "api.thedp.com: Invalid limit. Must be int"})
}

func TestLimitRequired(t *testing.T) {
	t.Logf("Testing GET /dp/articles")

	req, _ := http.NewRequest("GET", "/dp/articles", nil)
	res := executeRequest(t, req)

	assertStatus(t, res, http.StatusBadRequest)
	assertEqualJson(t, res, map[string]string{"message": "api.thedp.com: If no filter parameters are specified, limit must be specified"})
}

func TestGetByAuthors(t *testing.T) {
	t.Logf("Testing GET /dp/articles?author=jared-mitovich&author=sara-forastieri")

	req, _ := http.NewRequest("GET", "/dp/articles?author=jared-mitovich&author=sara-forastieri", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) >= 1 {
		t.Log("Expected >= 1 articles. Success")
	} else {
		t.Errorf("Expected >= 1 articles. Got %d\n", len(articles))
	}

	// TODO: ensure articles have correct authors
}

func TestGetByNonexistentAuthor(t *testing.T) {
	t.Logf("Testing GET /dp/articles?author=nonexistent-author")

	req, _ := http.NewRequest("GET", "/dp/articles?author=nonexistent-author", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) == 0 {
		t.Log("Expected 0 articles. Success")
	} else {
		t.Errorf("Expected 0 articles. Got %d\n", len(articles))
	}
}

func TestGetByTags(t *testing.T) {
	t.Logf("Testing GET /dp/articles?tag=news&tag=front")

	req, _ := http.NewRequest("GET", "/dp/articles?tag=news&tag=front", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) >= 1 {
		t.Log("Expected >= 1 articles. Success")
	} else {
		t.Errorf("Expected >= 1 articles. Got %d\n", len(articles))
	}

	// insure articles have tags
}

func TestGetByNonexistentTag(t *testing.T) {
	t.Logf("Testing GET /dp/articles?tag=nonexistent-tag")

	req, _ := http.NewRequest("GET", "/dp/articles?tag=nonexistent-tag", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) == 0 {
		t.Log("Expected 0 articles. Success")
	} else {
		t.Errorf("Expected 0 articles. Got %d\n", len(articles))
	}
}

func TestGetBySlug(t *testing.T) {
	t.Logf("Testing GET /dp/articles?slug=penn-flu-clinic-2023-recap-vaccines")

	req, _ := http.NewRequest("GET", "/dp/articles?slug=penn-flu-clinic-2023-recap-vaccines", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) == 1 {
		t.Log("Expected 1 articles. Success")
	} else {
		t.Errorf("Expected 1 articles. Got %d\n", len(articles))
	}

	if articles[0].Slug == "penn-flu-clinic-2023-recap-vaccines" {
		t.Log("Expected slug to be penn-flu-clinic-2023-recap-vaccines. Success")
	} else {
		t.Errorf("Expected slug to be penn-flu-clinic-2023-recap-vaccines. Got %s\n", articles[0].Slug)
	}
}
func TestGetByNonexistentSlug(t *testing.T) {
	t.Logf("Testing GET /dp/articles?slug=nonexistent-slug")

	req, _ := http.NewRequest("GET", "/dp/articles?slug=nonexistent-slug", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) == 0 {
		t.Log("Expected 0 articles. Success")
	} else {
		t.Errorf("Expected 0 articles. Got %d\n", len(articles))
	}
}
func TestSortPopular(t *testing.T) {
	t.Logf("Testing GET /dp/articles?sort=popular&limit=5")

	req, _ := http.NewRequest("GET", "/dp/articles?sort=popular&limit=5", nil)
	res := executeRequest(t, req)

	assertSuccess(t, res)
	articles := extractArticles(t, res)

	if len(articles) == 5 {
		t.Log("Expected 5 articles. Success")
	} else {
		t.Errorf("Expected 5 articles. Got %d\n", len(articles))
	}

	// TODO: ensure sorted properly
}

func executeRequest(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	t.Helper()

	rec := httptest.NewRecorder()
	a.Router.ServeHTTP(rec, req)

	return rec
}

func extractArticles(t *testing.T, res *httptest.ResponseRecorder) []api.Article {
	t.Helper()

	var articles []api.Article
	err := json.Unmarshal(res.Body.Bytes(), &articles)

	if err != nil {
		t.Errorf("Expected response to be an array of articles. Got %s\n", res.Body.String())
	} else {
		t.Log("Expected response to be an array of articles. Success")
	}

	return articles
}

func assertSuccess(t *testing.T, res *httptest.ResponseRecorder) {
	t.Helper()
	assertStatus(t, res, http.StatusOK)
}

func assertStatus(t *testing.T, res *httptest.ResponseRecorder, code int) {
	t.Helper()
	if res.Code != code {
		t.Errorf("Expected response code %d. Got %d\n", code, res.Code)
	} else {
		t.Logf("Expected response code %d. Success", code)
	}
}

func assertEqualJson(t *testing.T, res *httptest.ResponseRecorder, expected interface{}) {
	t.Helper()
	expectedJson, _ := json.Marshal(expected)
	if res.Body.String() != string(expectedJson) {
		t.Errorf("Expected response %s. Got %s\n", string(expectedJson), res.Body.String())
	} else {
		t.Logf("Expected response %s. Success", string(expectedJson))
	}
}
