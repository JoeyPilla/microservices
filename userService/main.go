package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "user-database"
	port     = 5483
	user     = "user"
	password = "user_password"
	dbname   = "user_db"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from the posts!")
}
