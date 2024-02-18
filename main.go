type Item struct {
	Name string `json:"name"`
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range items {
		if item.Name == params["name"] {
			items = append(items[:idx], items[idx+1:]...)
			var newItem Item
			_ = json.NewDecoder(r.Body).Decode(&newItem)
			items = append(items, newItem)
			json.NewEncoder(w).Encode(newItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range items {
		if item.Name == params["name"] {
			items = append(items[:idx], items[idx+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{name}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{name}", deleteItem).Methods("DELETE")
	http.ListenAndServe(":8000", r)
}