package web

import (
	"bytes"
	"container/list"
	"fmt"
	"go-learn/meal-coupon/constants"
	"go-learn/meal-coupon/model"
	"go-learn/meal-coupon/service"
	"net/http"
	"strconv"
)

func QueryCouponInfo(w http.ResponseWriter, r *http.Request) {
	loginName := GetLoginNameFromCookie(r)
	if loginName == nil {
		fmt.Fprint(w, "您还未登录")
	}
	staffInfo := service.GetStaffInfoByLoginName(loginName)
	if staffInfo == nil {
		fmt.Fprint(w, "登录信息不准确,请重新登录")
	}
	if staffInfo.Role != "admin" {
		teamInfo := service.GetDataByTeam(&staffInfo.TeamNo)
		fmt.Fprint(w, assemTeamData(teamInfo))
	} else {
		teamList := service.ListAllTeamInfo()

		fmt.Fprint(w, assemAllTeamData(teamList))
	}
}

func assemTeamData(teamInfo *model.Team) string {
	teamValue := bytes.Buffer{}
	teamValue.WriteString("您当前的组为：")
	teamValue.WriteString(teamInfo.Name)
	teamValue.WriteString(constants.NewLine)
	teamValue.WriteString("小组当前餐券为：")
	teamValue.WriteString(strconv.Itoa(teamInfo.CouponNum))
	return teamValue.String()
}

func assemAllTeamData(teamList *list.List) string {
	teamValue := bytes.Buffer{}

	for team := teamList.Front(); team != nil; team.Next() {
		teamInfo := team.Value.(model.Team)
		teamValue.WriteString("组名：")
		teamValue.WriteString(teamInfo.Name)
		teamValue.WriteString(" ,  餐券数量：")
		teamValue.WriteString(strconv.Itoa(teamInfo.CouponNum))
		teamValue.WriteString(constants.NewLine)
	}
	return teamValue.String()
}
