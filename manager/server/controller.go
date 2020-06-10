package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var BaseUrl = "http://localhost:"

func (s *Server) getWorkerAddress() (string, int) {
	min := 0
	pos := 0

	for indx, worker := range *s.Workers {
		if indx == 0 {
			min = worker.KeyCount
			pos = indx
		} else {
			if worker.KeyCount < min {
				min = worker.KeyCount
				pos = indx
			}
		}
	}

	(*s.Workers)[pos].Lock()
	(*s.Workers)[pos].KeyCount++
	(*s.Workers)[pos].Unlock()

	return (*s.Workers)[pos].Addr, (*s.Workers)[pos].Id
}

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check is key already present, error if yes
	_, ok := s.CacheIndex[req.Key]
	if ok {
		http.Error(w, "Key Already Present", http.StatusAlreadyReported)
		return
	}

	port, id := s.getWorkerAddress()
	fmt.Println("Using worker" + strconv.Itoa(id) + "Running on port" + port)

	s.CacheIndex[req.Key] = CacheIndexRecord{
		WorkerId:  id,
		Key:       req.Key,
		Addr:      port,
		CreatedAt: time.Now(),
	}

	workerURL := BaseUrl + port + "/add"
	reqBody, _ := json.Marshal(req)
	resFromWorker, err1 := http.Post(workerURL, "application/json", bytes.NewBuffer(reqBody))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	body, _ := ioutil.ReadAll(resFromWorker.Body)
	w.Write(body)
}

func (s *Server) RemoveFromCache(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://localhost:3001/remove")
	body, _ := ioutil.ReadAll(resp.Body)
	finalRes := append([]byte("Hello from Manager, Remove Api "), body...)
	w.Write(finalRes)
}

func (s *Server) GetFromCache(w http.ResponseWriter, r *http.Request) {
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, ok := s.CacheIndex[req.Key]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	port := record.Addr
	workerURL := BaseUrl + port + "/get"

	reqBody, _ := json.Marshal(req)
	resFromWorker, err1 := http.Post(workerURL, "application/json", bytes.NewBuffer(reqBody))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	body, _ := ioutil.ReadAll(resFromWorker.Body)
	w.Write(body)
}
