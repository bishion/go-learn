package base

import "fmt"

const (
	i      = 100
	pi     = 3.14
	prefix = "Go_"
)

var (
	company  string
	position string
)

const (
	x = iota
	y
	z       = "guo"
	d, e, f = iota, iota, iota
	g       = iota
)

func print() {
	fmt.Println(x, y, z, d, e, f, g)
}
