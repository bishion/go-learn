package service

import (
	"fmt"
	"go-learn/meal-coupon/constants"
	"go-learn/meal-coupon/model"
	"io/ioutil"
	"strconv"
	"strings"
)

/**
 * 数据服务, 启动时加载数据, 并提供修改参数功能
 */

var teamMap = make(map[string]*model.Team)
var staffMap = make(map[string]*model.Staff)

/**
 * 启动时加载数据
 * 1. 读取 team  信息, 放入 teamMap 中
 * 2. 读取 staff 信息, 放入 staffMap 中
 */
func init() {

	initTeamFile()
	initStaffFile()

}
func ModifyCoupon(teamNo string, num int) {
	team := teamMap[teamNo]
	team.CouponNum = num
	fmt.Println(teamMap["101"].CouponNum)
}

func initStaffFile() {
	staffFile, err := ioutil.ReadFile(constants.Staff_file)
	if err != nil {
		panic(err)
	}
	for index, line := range strings.Split(string(staffFile), "\n") {
		if index == 0 {
			continue
		}
		attr := strings.Split(line, ",")
		if err != nil {
			fmt.Println("数据中有非法字符:" + line)
			panic(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if len(attr) < 5 {
			panic("该数据有错误:" + line)
		}
		staffMap[strings.TrimSpace(attr[0])] = &model.Staff{strings.TrimSpace(attr[0]), strings.TrimSpace(attr[1]), strings.TrimSpace(attr[2]), strings.TrimSpace(attr[3]), strings.TrimSpace(attr[4])}
	}
}
func initTeamFile() {
	teamFile, err := ioutil.ReadFile(constants.Team_file)
	if err != nil {
		panic(err)
	}
	for index, line := range strings.Split(string(teamFile), "\n") {
		if index == 0 {
			continue
		}

		attr := strings.Split(line, ",")
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if len(attr) < 4 {
			panic("该数据有错误:" + line)
		}
		couponNum, err := strconv.Atoi(strings.TrimSpace(attr[3]))
		if err != nil {
			fmt.Println("数据中有非法字符:" + line)
			panic(err)
		}
		teamMap[strings.TrimSpace(attr[0])] = &model.Team{strings.TrimSpace(attr[0]), strings.TrimSpace(attr[1]), strings.TrimSpace(attr[2]), couponNum}
	}
}
