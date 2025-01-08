package handlers

import "net/http"

func OpenEP(w http.ResponseWriter, r *http.Request) {
	// OpenEP handler
	w.Write([]byte("This is a open endpoint"))
}

func ProtectedEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected endpoint"))
}
