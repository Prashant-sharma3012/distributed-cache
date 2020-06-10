package server

import (
	"net/http"
	"sync"
	"time"
)

type Worker struct {
	sync.Mutex
	Id       int
	KeyCount int
	Addr     string
}

type CacheIndexRecord struct {
	WorkerId  int
	Key       string
	CreatedAt time.Time
}

type Server struct {
	Srv        *http.Server
	Handler    *http.ServeMux
	Workers    *[]Worker
	CacheIndex *[]CacheIndexRecord
}

func InitServer(port string) *Server {
	handler := http.NewServeMux()

	return &Server{
		Srv: &http.Server{
			Addr:    port,
			Handler: handler,
		},
		Handler: handler,
	}
}
