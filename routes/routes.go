package routes

import (
	"net/http"

	"github.com/aramakam3505/hyperAPI/handlers"
)

func IntializeRoutes() {
	http.HandleFunc("/hello", handlers.Hello)
}
