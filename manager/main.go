package manager

import (
	"log"

	"github.com/distributed-cache/manager/server"
)

func main() {
	s := server.InitServer()
	log.Fatal(s.Srv.ListenAndServe())
}
