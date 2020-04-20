package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Function which start the server
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/hello/{name}", hello).Methods("GET")
	router.HandleFunc("/query", query).Methods("GET")
	router.HandleFunc("/post", post).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", router))
}

/* Write plain text */
/* curl localhost:4000 */
// index start
func index(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `Pong`)
}

/* Path arguments */
/* curl localhost:4000/hello/Dario */
func hello(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello " + vars["name"],
	}
	json.NewEncoder(w).Encode(response)
}

/* Query arguments */
/* curl localhost:4000/query?key=test */
func query(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")

	response := struct {
		Key string `json:"key"`
	}{
		Key: key,
	}
	json.NewEncoder(w).Encode(response)
}

/*
 curl localhost:4000/post -H "Content:application/json" --data '{"data":"test"}'
 curl localhost:4000/post -H "Content:application/json" --data '{"data":"asd","test":"test"}'
*/
func post(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	var request struct {
		Data *string
	}
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(request)
}
