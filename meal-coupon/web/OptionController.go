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
	"time"
)

func QueryCouponInfo(w http.ResponseWriter, r *http.Request) {
	loginName := GetLoginNameFromCookie(r)
	if loginName == nil {
		fmt.Fprint(w, "您还未登录")
		return
	}
	staffInfo := service.GetStaffInfoByLoginName(loginName)
	if staffInfo == nil {
		fmt.Fprint(w, "登录信息不准确,请重新登录")
		return
	}
	if staffInfo.Role != constants.RoleAdmin {
		teamInfo := service.GetDataByTeam(&staffInfo.TeamNo)
		fmt.Fprint(w, assemTeamData(teamInfo))
	} else {
		teamList := service.ListAllTeamInfo()

		fmt.Fprint(w, assemAllTeamData(teamList))
	}
}
func ModifyCoupon(w http.ResponseWriter, r *http.Request) {

	loginName := GetLoginNameFromCookie(r)
	if loginName == nil {
		fmt.Fprint(w, "您还未登录")
		return
	}
	staffInfo := service.GetStaffInfoByLoginName(loginName)
	if staffInfo == nil {
		fmt.Fprint(w, "登录信息不准确,请重新登录")
		return
	}

	if staffInfo.Role != constants.RoleLeader {
		fmt.Fprint(w, "您不是组长，没有权限修改餐券")
		return
	}
	values := r.URL.Query()

	couponNumStr := values["couponNum"]
	if len(couponNumStr) == 0 {
		fmt.Fprint(w, "您还未输入餐券数量")
		return
	}

	couponNum, err := strconv.Atoi(couponNumStr[0])
	if err != nil {
		fmt.Fprint(w, "请输入正确的数值类型："+couponNumStr[0])
		return
	}
	today := time.Now().Day()
	if today > today {
		fmt.Fprint(w, "只有每个月 5 号之前才能修改！")
		return
	}
	service.ModifyCoupon(staffInfo.TeamNo, couponNum)
	fmt.Fprint(w, "操作成功，当前餐券数为："+couponNumStr[0])

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

	for team := teamList.Front(); team != nil; team = team.Next() {
		teamInfo := team.Value.(model.Team)
		teamValue.WriteString("组名：")
		teamValue.WriteString(teamInfo.Name)
		teamValue.WriteString(" ,  餐券数量：")
		teamValue.WriteString(strconv.Itoa(teamInfo.CouponNum))
		teamValue.WriteString(constants.NewLine)
	}
	return teamValue.String()
}
