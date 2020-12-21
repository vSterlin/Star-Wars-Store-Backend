package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vSterlin/sw-store/internal/app/store"
)

func GetOneHandler(pr *store.ProductRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			panic(err)
		}

		product, err := pr.FindById(id)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}
