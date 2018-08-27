package base

import "fmt"

// array
var arr [10]int

func print_arr() {
	arr[0] = 1
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6}
	fmt.Println(a[1], b[9], c)

}

func print_2d_arr() {
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArray)
}

// slice
var fslice []int
var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

func print_slice() {
	slice := []byte{'a', 'b'}
	fmt.Println(slice)

	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]

	fmt.Printf(string(b) + string(a))

	a = ar[:3]
	a = ar[5:]
	a = ar[:]

	b = ar[3:7]
	b = a[1:3]
	b = a[:3]
	b = a[0:5]
	b = a[:]

	_ = append(a, 'c')

	b = ar[2:4:10]

}

// map
var numbers map[string]int
var numbers2 = make(map[string]int)

func print_map() {
	numbers = make(map[string]int)

	numbers["one"] = 1
	numbers["two"] = 1
	numbers["three"] = 1
	fmt.Printf("第三个数字是：%d", numbers["three"])

	delete(numbers, "two")
	fmt.Println(numbers["two"])
}
