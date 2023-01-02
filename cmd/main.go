package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sour49/bookingsystem/pkg/routes"
)

const portNumber = ":8888"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
