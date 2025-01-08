package routes

import (
	"authservice/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Open EP
	r.HandleFunc("/open", handlers.OpenEP).Methods("GET")

	// Login EP for generating JWT
	r.HandleFunc("/login", handlers.LoginEP).Methods("POST")

	// Protected EP, uses JWT for authentication
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(JWTMiddleware)

	protected.HandleFunc("", handlers.ProtectedEP).Methods("GET")

	return r
}
