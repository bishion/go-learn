package constants

import "os"

var (
	baseDir, _    = os.Getwd()
	workDir       = baseDir + "/data/"
	StaffFile     = workDir + "staff.dat"
	TeamFile      = workDir + "team.dat"
	TeamDataTitle = "组号, 组名, 组长, 现有餐券数量\n"
	NewLine       = "\n"
	Comma         = ","
	LoginName     = "loginName"
	RoleLeader    = "leader"
	RoleAdmin     = "admin"
)
