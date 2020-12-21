package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vSterlin/sw-store/internal/app/store"
)

func GetAllHandler(pr *store.ProductRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := pr.FindAll()
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(products)
	}
}
