package main

import "go-learn/meal-coupon/service"

/**
*   加班餐券管理系统
•	已有整个部门的人员名单和各team的人员名单，和每个team的leader和部门admin的人员信息
•	可以通过员工号登陆
•	team member登陆后能看到自己team的当月的餐券申请数量，
•	team leader登陆后不仅可以看到数量，每月5号之前还可以编辑餐券数量
•	admin登陆后可以看到各个team的申请数量
*/
func main() {

	service.ModifyCoupon("101", 12)
}
