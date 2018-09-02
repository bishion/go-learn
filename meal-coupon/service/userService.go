package service

import (
	"go-learn/meal-coupon/model"
	"errors"
)

func Login(username *string, password *string) (*model.Staff,error) {
	staffInfo := staffMap[*username]
	if staffInfo == nil || *password != staffInfo.Password {
		return nil, errors.New("用户名或密码错误")
	}
	return staffInfo, nil
}
