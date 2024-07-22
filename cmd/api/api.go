package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hilalerdogan/mobile-application-backend/service/user"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{address: address, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Println("Server is running on", s.address)
	return http.ListenAndServe(s.address, router)
}
