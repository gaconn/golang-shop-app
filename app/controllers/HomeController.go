package controllers

import (
	"net/http"

	"github.com/quan12xz/golang-shop-app/utils"
)

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	utils.DBConnect()
}
