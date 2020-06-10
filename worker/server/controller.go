package server

import (
	"net/http"
	"strconv"
)

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Key Added to Wroker" + strconv.Itoa(s.Id)))
}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Key Removed From Wroker" + strconv.Itoa(s.Id)))
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {

}
