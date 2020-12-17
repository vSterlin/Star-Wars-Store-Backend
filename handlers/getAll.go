package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vSterlin/sw-store/repos"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {

	pr := repos.NewProductRepo()
	products := pr.FindAll()
	fmt.Printf("%+v", products)
	json.NewEncoder(w).Encode(products)
}
