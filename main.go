package main

import (
	"net/http"

	"github.com/aramakam3505/hyperAPI/routes"
)

func main() {

	routes.IntializeRoutes()
	http.ListenAndServe(":8090", nil)
}
