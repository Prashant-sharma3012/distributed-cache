package manager

import (
	"fmt"
	"log"

	"github.com/distributed-cache/manager/api"
	"github.com/distributed-cache/manager/server"
)

func StartCacheServer() {
	fmt.Println("Starting Manager")
	s := server.InitServer()

	fmt.Println("Initializing End Points")
	api.InitRoutes(s)

	fmt.Println("Initialized Server, listening on post 3000")
	log.Fatal(s.Srv.ListenAndServe())
}
