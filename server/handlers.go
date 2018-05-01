package server

import (
	"cources/homework/models"
	"cources/homework/my_db"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Banco de España!\n")
}

func ClientIndex(w http.ResponseWriter, r *http.Request) {
	clients := my_db.FindClients()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clients); err != nil {
		panic(err)
	}
}

func ClientShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	clients := my_db.FindClient(params["id"])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clients); err != nil {
		panic(err)
	}
}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"id":"2","name":"i","email":"i@epam.com","phone":"202","sms_notification":"true","email_notification":"true"}' http://localhost:8080/clients

*/

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &client); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	c_client := my_db.CreateClient(client)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c_client); err != nil {
		panic(err)
	}
}
