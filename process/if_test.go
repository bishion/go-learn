package process

import "fmt"

var age int

func if_test() {
	// 简单的 if ... else ...
	if age > 18 {
		fmt.Println("I am adult")
	} else {
		fmt.Println("I am child")
	}
	// if 中定义变量
	if x := getAge(); x > 18 {
		fmt.Print("I am adult")
	} else {
		fmt.Print("I am child")
	}
	// if else if

	if x := getAge(); x > 18 {
		fmt.Print("I am adult")
	} else if x > 12 {
		fmt.Print("I am teenager")
	} else {
		fmt.Print("I am child")
	}

}
func getAge() int {
	return 18
}
