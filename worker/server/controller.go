package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Cache[req.Key] = req
	w.Write([]byte("Key Added to Wroker" + strconv.Itoa(s.Id)))
}

func (s *Server) ReplaceInCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Remove record from server
	if req.KeyToDelete != "" {
		fmt.Println("Removing Key From cache: " + req.KeyToDelete)
		delete(s.Cache, req.KeyToDelete)
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
		return
	}

	val, ok := s.Cache[req.Key]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(val)
}
