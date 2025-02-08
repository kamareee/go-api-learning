package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Id   string `json:"id"`
	Item string `json:"item"`
}

type ItemId struct {
	Id string `json:"id"`
}

func main() {

	//var todos = make([]TodoItem, 0)
	var todos = make(map[string]string)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
		return
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, request *http.Request) {
		var t TodoItem
		err := json.NewDecoder(request.Body).Decode(&t)

		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		todos[t.Id] = t.Item
		writer.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("DELETE /todo", func(w http.ResponseWriter, request *http.Request) {
		var item ItemId
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			log.Println(err)
		}
		delete(todos, item.Id)

		_, err = w.Write([]byte("deleted " + item.Id))
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
