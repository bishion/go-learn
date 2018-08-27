package base

import "fmt"

// 定义变量并指定类型
var name string

// 同时定义多个变量
var age, score int8

// 定义变量并赋值
var birthday string = "1989-11-08"

// 同时定义多个变量并赋值
var name1, name2, name3 string = "guo", "bizi", "bishion"

// 同时定义多个变量并赋值, 可以省略变量类型
var age1, age2, age3 = 18, 28, 30

var _ = "赋值给下划线的变量都会被丢弃"

func method() {
	// 最简单的定义方式，不过需要放在方法里
	score1, score2, score3 := 90, 80, 70
	// 变量定义了就要使用, 不然编译不通过
	fmt.Print(score1)
	fmt.Print(score2)
	fmt.Print(score3)
}
