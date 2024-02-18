package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

func newRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/items/count", getItemCount).Methods("GET")
    return r
}

func TestGetItemCount(t *testing.T) {
    req, err := http.NewRequest("GET", "/items/count", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    router := newRouter()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var itemCount int
    err = json.Unmarshal(rr.Body.Bytes(), &itemCount)
    if err != nil {
        t.Fatal(err)
    }

    if itemCount != 0 {
        t.Errorf("Handler returned unexpected body: got %v want %v", itemCount, 0)
    }
}