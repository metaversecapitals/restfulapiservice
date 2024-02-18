package main 

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var items []interface{}

type ItemCountHandler struct{}

func (h ItemCountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(len(items))
}

func main() {
	r := mux.NewRouter()
	r.Handle("/items/count", ItemCountHandler{}).Methods(http.MethodGet)
	http.ListenAndServe(":8000", r)
}