package web

import (
	"go-learn/meal-coupon/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.HasNil()
}
