package server

import (
	"net/http"
)

type Req struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Server struct {
	Id      int
	Srv     *http.Server
	Handler *http.ServeMux
	Cache   map[string]interface{}
}

func InitServer(port string, id int) *Server {
	handler := http.NewServeMux()

	return &Server{
		Srv: &http.Server{
			Addr:    port,
			Handler: handler,
		},
		Handler: handler,
		Id:      id,
		Cache:   make(map[string]interface{}),
	}
}
