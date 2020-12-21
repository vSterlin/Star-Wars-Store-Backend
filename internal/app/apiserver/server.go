package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/vSterlin/sw-store/internal/app/handlers"
	"github.com/vSterlin/sw-store/internal/app/store"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  *store.Store
}

func newServer(store *store.Store) *server {

	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	pr := s.store.Product()
	s.router.HandleFunc("/products", handlers.GetAllHandler(pr)).Methods("GET")
	s.router.HandleFunc("/products/{id}", handlers.GetOneHandler(pr)).Methods("GET")

}
