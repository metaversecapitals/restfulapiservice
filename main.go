// A function that gets the total count of items.
func getItemCount(w http.ResponseWriter, r *http.Request) {
    // Setting the content type of the response to JSON.
    w.Header().Set("Content-Type", "application/json")
    // Returning the count of all items in the JSON format.
    json.NewEncoder(w).Encode(len(items))
}

// Add this line to the main function to handle the new endpoint.
r.HandleFunc("/items/count", getItemCount).Methods("GET")