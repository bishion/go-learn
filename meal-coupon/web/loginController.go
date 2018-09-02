package web

import (
	"fmt"
	"go-learn/meal-coupon/constants"
	"go-learn/meal-coupon/model"
	"go-learn/meal-coupon/service"
	"go-learn/meal-coupon/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	defer func(){
		if err := recover();err!=nil{
			fmt.Fprint(w, err)
			return
		}
	}()
	values := r.URL.Query()

	username, password := values["username"], values["password"]
	if len(username) == 0 || len(password) == 0 {
		fmt.Fprint(w, "用户名或密码为空")
		return
	}
	staffInfo,err := service.Login(&username[0], &password[0])
	if err != nil{
		fmt.Fprint(w, err)
		return
	}

	// 设置 cookie, 这里使用cookie的方式保存登录信息,并不安全
	// 然而 go 并原生支持 session, 实现起来比较麻烦
	// 这里就将就用 cookie 了
	setLoginName2Cookie(w, staffInfo)
	fmt.Fprint(w, "登陆成功！")
}

func setLoginName2Cookie(w http.ResponseWriter, staffInfo *model.Staff) {
	loginCookie := http.Cookie{
		Name:     constants.LoginName,
		Value:    staffInfo.No,
		HttpOnly: true,
	}
	http.SetCookie(w, &loginCookie)
}

func GetLoginNameFromCookie(r *http.Request) *string {
	loginCookie, err := r.Cookie(constants.LoginName)
	utils.CheckNil(err)
	return &loginCookie.Value
}
