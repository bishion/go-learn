package base

import (
	"errors"
	"fmt"
)

// 布尔类型, 默认是 false
var isActive bool

var enabled, disabled = true, false

func test() {
	var available bool
	valid := true
	available = false
	print(valid)
	print(available)
}

// 数值类型
var a int8
var b int32

func numerical() {
	c := a + 1
	print(c)
}

// 复数
var c complex64 = 5 + 5i

func print_complex() {
	fmt.Printf("Value is : %v", c)
}

// 字符串
var s string = "hello"

func pring_string() {
	//s[0] = 'c' // sting 类型不能修改, 这里编译报错
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	s3 := s + s2
	s4 := "c" + s[1:]

	fmt.Printf("s3:%s, s4:%s", s3, s4)

	m := `hello 
                world`
	fmt.Printf("m:%s", m)
}

// error
func print_error() {
	err := errors.New("I am an error")
	if err != nil {
		fmt.Print(err)
	}
}
