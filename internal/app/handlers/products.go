package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vSterlin/sw-store/repos"
)

func GetOneHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	pr := repos.NewProductRepo()
	product, err := pr.FindOne(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {

	pr := repos.NewProductRepo()
	products := pr.FindAll()
	fmt.Printf("%+v", products)
	json.NewEncoder(w).Encode(products)
}
