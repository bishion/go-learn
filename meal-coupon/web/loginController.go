package web

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
}
