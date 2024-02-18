type Item struct {
	Name string `json:"name"`
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.Handle("/items/count", ItemCountHandler{}).Methods(http.MethodGet)
	http.ListenAndServe(":8000", r)
}