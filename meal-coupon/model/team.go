package model

/**
 * 组信息: 组号, 组名, 组长, 现有餐券数量
 */
type Team struct {
	No        string
	Name      string
	Leader    string
	CouponNum int
}
