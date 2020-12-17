package main

import (
	"net/http"

	"github.com/vSterlin/sw-store/database"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/vSterlin/sw-store/handlers"
)

func main() {
	defer database.DB.Close()
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetHandler)
	http.ListenAndServe("localhost:8080", r)

}
