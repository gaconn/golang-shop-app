package routes

import (
	"net/http"

	"github.com/quan12xz/golang-shop-app/controllers"
)

func RoutesHandler() {
	http.HandleFunc("/", controllers.CheckAuth)
}
