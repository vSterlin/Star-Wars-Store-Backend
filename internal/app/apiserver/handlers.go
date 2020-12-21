package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) getOneHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			panic(err)
		}

		product, err := s.store.Product().FindById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}

func (s *server) getAllHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, _ := s.store.Product().FindAll()
		fmt.Printf("%+v", products)
		json.NewEncoder(w).Encode(products)
	}
}
