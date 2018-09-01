package utils

func CheckNil(err error) {
	if err != nil {
		panic(err)
	}
}
