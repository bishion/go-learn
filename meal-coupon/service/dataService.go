package service

import (
	"bytes"
	"container/list"
	"fmt"
	"go-learn/meal-coupon/constants"
	"go-learn/meal-coupon/model"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

/**
 * 数据服务, 启动时加载数据, 并提供数据查询和修改参数功能
 */

var teamMap = make(map[string]*model.Team)
var staffMap = make(map[string]*model.Staff)
var mutex sync.Mutex

/**
 * 启动时加载数据
 * 1. 读取 team  信息, 放入 teamMap 中
 * 2. 读取 staff 信息, 放入 staffMap 中
 */
func init() {
	initTeamFile()
	initStaffFile()

}

// 数据修改功能, 提供组号和修改后的值. 修改完毕后直接持久化到文件
func ModifyCoupon(teamNo string, num int) {
	team := teamMap[teamNo]
	team.CouponNum = num
	fmt.Println(teamMap["101"].CouponNum)
	writeTeamFile()
}

func GetDataByTeam(teamNoPtr *string) *model.Team {
	return teamMap[*teamNoPtr]
}

func ListAllTeamInfo() *list.List {
	teamList := list.New()
	for _, value := range teamMap {
		teamList.PushBack(*value)
	}
	return teamList
}
func GetStaffInfoByLoginName(staffNo *string) *model.Staff {
	return staffMap[*staffNo]
}

// 将小组数据写入文件. 这里为了方便起见, 直接统一全量写入文件.
// 为避免多线程同时读取文件,这里使用单线程
func writeTeamFile() {
	mutex.Lock()

	teamValue := bytes.Buffer{}
	teamValue.WriteString(constants.TeamDataTitle)
	for _, value := range teamMap {
		teamValue.WriteString(value.No)
		teamValue.WriteString(constants.Comma)
		teamValue.WriteString(value.Name)
		teamValue.WriteString(constants.Comma)
		teamValue.WriteString(value.Leader)
		teamValue.WriteString(constants.Comma)
		teamValue.WriteString(strconv.Itoa(value.CouponNum))
		teamValue.WriteString(constants.NewLine)
	}

	err := ioutil.WriteFile(constants.TeamFile, []byte(teamValue.String()), 0666)
	if err != nil {
		panic(err)
	}

	mutex.Unlock()
}
func initStaffFile() {
	staffFile, err := ioutil.ReadFile(constants.StaffFile)
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
	teamFile, err := ioutil.ReadFile(constants.TeamFile)
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
