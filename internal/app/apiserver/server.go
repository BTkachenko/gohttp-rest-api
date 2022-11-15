package apiserver

import (
	"net/http"

	"github.com/BTkachenko/gotest/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store  //interface
}

//Przyjmuje interface schroniska 
func newServer(store store.Store) *server{
	s := &server{
		router:mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users",s.handleUsersCreate()).Methods("Post")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {

	}
}


func (s *server) ServeHTTP( w http.ResponseWriter, r *http.Request){
	s.router.ServeHTTP(w,r)
}