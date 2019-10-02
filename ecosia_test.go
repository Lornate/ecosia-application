package main

import (
	"ecosia-application/app/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFavouriteTree(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/?favoriteTree=apple", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "Please tell me your favorite tree" {
		t.Errorf("Expected It's nice to know that your favorite tree is a apple Got %s", body)
	}
}

func TestGetFavouriteTreeWithoutQueryParameters(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/?", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	if body := response.Body.String(); body != "Please tell me your favorite tree\n" {
		t.Errorf("Expected Please tell me your favorite tree Got %s", body)
	}
}

func TestGetFavouriteTreeWithMalfunctionedQueryParameters(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/?favorite=apple", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
	if body := response.Body.String(); body != "Please tell me your favorite tree\n" {
		t.Errorf("Expected Please tell me your favorite tree Got %s", body)
	}
}

func TestGetFavouriteTreeInvalidRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8000/api", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	if body := response.Body.String(); body == "" {
		t.Errorf(" Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.GetRoutes().ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
