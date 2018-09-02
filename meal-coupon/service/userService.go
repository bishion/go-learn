package service

import "go-learn/meal-coupon/model"

func Login(username *string, password *string) *model.Staff {
	staffInfo := staffMap[*username]
	if staffInfo == nil || *password != staffInfo.Password {
		panic("用户名或密码错误")
	}
	return staffInfo
}
