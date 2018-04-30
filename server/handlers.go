package server

import (
	_ "cources/homework/models"
	"cources/homework/my_db"
	_ "encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"net/http"
	_ "strconv"

	_ "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Banco de España!\n")
}

func ClientIndex(w http.ResponseWriter, r *http.Request) {
	q := fmt.Sprintf("SELECT * FROM clients")
	clients := my_db.Find(q, "clients")
	fmt.Println(clients)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// if err := json.NewEncoder(w).Encode(models.clients); err != nil {
	// 	panic(err)
	// }
}
