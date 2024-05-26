package care

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// 函数可以有多个形参和多个返回值，返回值也可以和形参一样拥有参数名称
func SumCare(a, b int) (sum int, err error) {
	//user.Name = "nick"
	//user.Age = 18
	if a <= 0 && b <= 0 {
		err = errors.New("两个数相加不能同时小于0")
		return
	}
	fmt.Println("Hello OK!!!")
	sum = a + b
	return
}

// 值传递和引用传递
func ReferenceCare(a int, b *int) {
	a += 1
	*b += 1
}

// 全局变量
var g int
var G int

// 变量作用域
func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g
}
