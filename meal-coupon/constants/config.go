package constants

import "os"

var (
	base_dir, _ = os.Getwd()
	work_dir    = base_dir + "/data/"
	Staff_file  = work_dir + "staff.dat"
	Team_file   = work_dir + "team.dat"
)
