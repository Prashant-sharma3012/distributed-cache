package server

import (
	"net/http"
	"time"
)

type CacheIndex struct {
	AtNode    string
	Key       string
	CreatedAt time.Time
}

type Server struct {
	Srv        *http.Server
	Handler    *http.ServeMux
	CacheIndex *CacheIndex
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
