package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

var articles []Article

func viewHandler(response http.ResponseWriter, request *http.Request) {

	header := response.Header()
	header.Add("Content-Type", "application/json")

	params := mux.Vars(request)
	for _, item := range articles {
		if item.Id == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}

	json.NewEncoder(response).Encode(&Article{})
}

func main() {
	articles = append(articles, Article{"1", "Some Title 1", "This is a test description"})
	articles = append(articles, Article{"2", "My Title 2", "A really interesting article"})

	router := mux.NewRouter()
	router.HandleFunc("/{id}", viewHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
