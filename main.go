package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=v dbname=sw-store sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	query, err := db.Query("SELECT id, name, price FROM products;")
	if err != nil {

		panic(err)
	}
	type product struct {
		id    int
		name  string
		price float64
	}
	arr := []*product{}

	for query.Next() {
		prod := &product{}
		err = query.Scan(&prod.id, &prod.name, &prod.price)
		if err != nil {
			panic(err)
		}
		arr = append(arr, prod)

	}

	for i := range arr {
		fmt.Println(arr[i])
	}
	// r := mux.NewRouter()
	// r.HandleFunc("/", handlers.GetHandler)
	// http.ListenAndServe("localhost:8080", r)

}
