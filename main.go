
```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Item struct {
    ID          string `json:"id"`
    Description string `json:"description"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for _, item := range items {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
	        return
        }
    }
}

func createItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newItem Item
    _ = json.NewDecoder(r.Body).Decode(&newItem)

    items = append(items, newItem)
    json.NewEncoder(w).Encode(newItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, item := range items {
        if item.ID == params["id"] {
            items = append(items[:index], items[index+1:]...)

            var updated Item
            _ = json.NewDecoder(r.Body).Decode(&updated)
            items = append(items, updated)

            json.NewEncoder(w).Encode(updated)
	        return
        }
    }
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

	for index, item := range items {
        if item.ID == params["id"] {
            items = append(items[:index], items[index+1:]...)
            break
        }
    }

	json.NewEncoder(w).Encode(items)
}

func main() {
    r := mux.NewRouter()

    items = append(items, Item{ID: "1", Description: "Item One"})

    r.HandleFunc("/items", getItems).Methods("GET")
    r.HandleFunc("/items/{id}", getItem).Methods("GET")
    r.HandleFunc("/items", createItem).Methods("POST")
    r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}