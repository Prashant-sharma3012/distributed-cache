package server

import "net/http"

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Manager, Add Api"))
}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Manager, Remove Api"))
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Manager, Get Api"))
}
