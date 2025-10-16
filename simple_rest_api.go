package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{
	{1, "Pen", 10},
	{2, "Notebook", 50},
}
var nextID = 3

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = nextID
	nextID++
	items = append(items, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			fmt.Fprintf(w, "✅ Item %d deleted", id)
			return
		}
	}
	http.Error(w, "Item not found", 404)
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getItems(w, r)
		case "POST":
			addItem(w, r)
		case "DELETE":
			deleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", 405)
		}
	})

	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
