package manager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/distributed-cache/manager/api"
	"github.com/distributed-cache/manager/server"
	worker "github.com/distributed-cache/worker"
)

func StartCacheServer(serverPortPtr int, numOfWorkers int, workerPortStartAtPtr int) {
	fmt.Println("Starting Manager")
	serverPortStr := strconv.Itoa(serverPortPtr)

	s := server.InitServer(":" + serverPortStr)

	fmt.Println("Initializing End Points")
	api.InitRoutes(s)

	fmt.Println("Start Workers")
	for i := 0; i < numOfWorkers; i++ {
		workerPortStr := strconv.Itoa(workerPortStartAtPtr + i)
		fmt.Println("Starting Worker at " + workerPortStr)
		go func() { worker.StartWorker(":" + workerPortStr) }()
	}

	fmt.Println("Initialized Server, listening on post 3000")
	log.Fatal(s.Srv.ListenAndServe())
}
