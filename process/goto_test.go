package process

func testGoto() int {
	i := 1
Here:
	print(i)
	i++
	if i < 5 {
		goto Here
	}
	return i
}
