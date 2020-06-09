package worker

import (
	"fmt"
	"log"

	"github.com/distributed-cache/worker/api"
	"github.com/distributed-cache/worker/server"
)

func StartWorker(port string) {
	fmt.Println("Starting Worker")
	s := server.InitServer(port)

	fmt.Println("Initializing End Points")
	api.InitRoutes(s)

	fmt.Println("Initialized Worker, listening on" + port)
	log.Fatal(s.Srv.ListenAndServe())
}
