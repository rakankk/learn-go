package main

import "fmt"

func main() {
	var fruit string = "6 apples"
	var watermallan bool = false // true or false
	if watermallan {             //如果不满足条件，下行代码则不执行
		fruit = "1 apples"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy", fruit)

}
