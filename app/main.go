package main

import (
	"fmt"
	"net/http"

	"github.com/quan12xz/golang-shop-app/routes"
)

func main() {
	fmt.Println("Connecting to server...")
	routes.RoutesHandler()
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err)
	}
}
