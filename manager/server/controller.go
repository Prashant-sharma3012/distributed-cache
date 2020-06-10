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
	var workerToUse Worker
	min := 0

	for indx, worker := range *s.Workers {
		if indx == 0 {
			min = worker.KeyCount
			workerToUse = worker
		} else {
			if worker.KeyCount < min {
				min = worker.KeyCount
				workerToUse = worker
			}
		}
	}

	workerToUse.Lock()
	workerToUse.KeyCount++
	workerToUse.Unlock()

	return workerToUse.Addr, workerToUse.Id
}

func (s *Server) AddToCache(w http.ResponseWriter, r *http.Request) {
	port, id := s.getWorkerAddress()

	fmt.Println("Using worker" + strconv.Itoa(id) + "Running on port" + port)

	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

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
	}

	record, ok := s.CacheIndex[req.Key]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
	}

	port := record.Addr
	workerURL := BaseUrl + port + "/get"

	reqBody, _ := json.Marshal(req)
	resFromWorker, err1 := http.Post(workerURL, "application/json", bytes.NewBuffer(reqBody))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	body, _ := ioutil.ReadAll(resFromWorker.Body)
	w.Write(body)
}
