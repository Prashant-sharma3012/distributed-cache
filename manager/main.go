package manager

import (
	"fmt"
	"log"

	"github.com/distributed-cache/manager/api"
	"github.com/distributed-cache/manager/server"
	worker "github.com/distributed-cache/worker"
)

func StartCacheServer() {
	fmt.Println("Starting Manager")
	s := server.InitServer()

	fmt.Println("Initializing End Points")
	api.InitRoutes(s)

	fmt.Println("Start Workers")
	go func() { worker.StartWorker(":3001") }()

	fmt.Println("Initialized Server, listening on post 3000")
	log.Fatal(s.Srv.ListenAndServe())
}
