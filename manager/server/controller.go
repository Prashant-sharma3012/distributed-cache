package server

import (
	"io/ioutil"
	"net/http"
)

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3001/add")
	body, _ := ioutil.ReadAll(resp.Body)
	finalRes := append([]byte("Hello from Manager, Add Api "), body...)
	w.Write(finalRes)

}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3001/remove")
	body, _ := ioutil.ReadAll(resp.Body)
	finalRes := append([]byte("Hello from Manager, Remove Api "), body...)
	w.Write(finalRes)
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3001/get")
	body, _ := ioutil.ReadAll(resp.Body)
	finalRes := append([]byte("Hello from Manager, Get Api "), body...)
	w.Write(finalRes)
}
