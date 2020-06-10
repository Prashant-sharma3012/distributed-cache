package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	s.Cache[req.Key] = req
	w.Write([]byte("Key Added to Wroker" + strconv.Itoa(s.Id)))
}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Key Removed From Wroker" + strconv.Itoa(s.Id)))
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	val, ok := s.Cache[req.Key]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(val)
}
