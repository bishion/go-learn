package model

/**
 * 人员信息: 编号, 姓名, 密码, 所在组号, 角色(member, leader, admin)
 */
type Staff struct {
	No       string
	Name     string
	Password string
	TeamNo   string
	Role     string
}
