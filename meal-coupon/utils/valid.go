package utils

func HasNil(valArr ...interface{}) {
	for _, value := range valArr {
		if value == nil {
			panic("参数中有空值")
		}
	}
}
