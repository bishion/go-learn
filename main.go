package main

import (
	"fmt"
	"go-learn/meal-coupon/web"
	"log"
	"net/http"
)

/**
* 加班餐券管理系统
* 程序启动主类
 */
func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", web.Login)
	http.HandleFunc("/query", web.QueryCouponInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("应用程序启动错误.", err)
	}
	//service.ModifyCoupon("101", 1)
	//fmt.Println(service.ListAllTeamInfo().Back())
	//teamString := "101"
	//fmt.Println(service.GetDataByTeam(&teamString))
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
