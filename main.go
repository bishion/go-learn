package main

import (
	"fmt"
	"go-learn/meal-coupon/web"
	"log"
	"net/http"
)

/**
 * 加班餐券管理系统
 * 程序启动主类, 对外暴露三个接口：
 * 1. /login 登陆接口, 通过 get 方式将用户名密码传入(原谅我的不严谨)
 * 2. /query 查询接口, 管理员返回的是所有组的餐券信息, 非管理员返回的是当前组的餐券信息
 * 3. /modify 修改接口, 只有组长有权限修改自己组的餐券数量, 并且是在每个月的 5 号前才能修改
 */
func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", web.Login)
	http.HandleFunc("/query", web.QueryCouponInfo)
	http.HandleFunc("/modify", web.ModifyCoupon)
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
