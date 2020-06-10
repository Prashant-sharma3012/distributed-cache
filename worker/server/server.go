package server

import (
	"net/http"
)

type Cache struct {
	Data map[string]interface{}
}

type Server struct {
	Id int
	Srv     *http.Server
	Handler *http.ServeMux
	Cache   *Cache
}

func InitServer(port string, id int) *Server {
	handler := http.NewServeMux()

	return &Server{
		Srv: &http.Server{
			Addr:    port,
			Handler: handler,
		},
		Handler: handler,
		Cache:   &Cache{},
		Id: id,
	}
}
