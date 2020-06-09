package server

import "net/http"

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Worker, Add Api"))
}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Worker, Remove Api"))
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Worker, Get Api"))
}
