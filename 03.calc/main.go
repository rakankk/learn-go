package main

import "fmt"

func main() {
	var a, b int8 = 30, 10
	fmt.Println("加法：a + b = ", a+b)
	fmt.Println("减法：a - b = ", a-b)
	fmt.Println("乘法：a * b = ", a*b)
	fmt.Println("除法：a / b = ", a/b)
	fmt.Println("取余：a % b = ", a%b)

	fmt.Println(true && false == false)
}
