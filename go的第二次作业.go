// go第二次作业
package main

import (
	"fmt"
)

func jia(a int, b int) int {
	return a + b
}
func jian(a int, b int) int {
	return a - b
}
func cheng(a int, b int) int {
	return a * b
}
func chu(a int, b int) int {
	return a / b
}
func rty(a int, b int, qwe func(int, int) int) int { //可以调用一类函数
	return qwe(a, b)
}

func main() {
	a := 0
	b := 0
	var d string
	fmt.Scanf("%d  %s %d \n", &a, &d, &b)
	switch d {
	case "+":
		fmt.Println(rty(a, b, jia))
	case "-":
		fmt.Println(rty(a, b, jian))
	case "*":
		fmt.Println(rty(a, b, cheng))
	case "/":
		fmt.Println(rty(a, b, chu))
	}

}

/*
	package main

import (

	"fmt"
	"math"

)

	func ddd(a float64) float64 {
		r := a
		var x = r * r
		return x
	}

	func main() {
		f := math.Pi
		g := f * ddd(2)
		fmt.Println(g)
*/
