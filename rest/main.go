package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kokiebisu/go/rest/structure"
)

// Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data.
// Use json.Unmarshal if you already have the JSON data in memory.

// Comment
var Articles []structure.Article

func createNewArticle(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("called")
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var article structure.Article

	// jsonbyte -> struct
	json.Unmarshal(result, &article)

	Articles = append(Articles, article)

	json.NewEncoder(rw).Encode(article)
}

func returnAllArticles(rw http.ResponseWriter, r *http.Request) {
	// job of encoding our articles array into a JSON string and then writing as part of our response
	json.NewEncoder(rw).Encode(Articles)
}

func returnSingleArticle(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(rw).Encode(article)
		}
	}
}

func deleteSingleArticle(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	for i, article := range Articles {
		if id == article.ID {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
	fmt.Println("deleted")
}

func updateSingleArticle(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	var updatedArticle structure.Article
	err = json.NewDecoder(r.Body).Decode(&updatedArticle)
	for i, article := range Articles {
		if id == article.ID {
			Articles[i] = updatedArticle
		}
	}
	json.NewEncoder(rw).Encode(Articles)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", deleteSingleArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateSingleArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = append(Articles, structure.Article{ID: 1, Title: "first title", Desc: "first description", Content: "first content"})
	handleRequests()
}
