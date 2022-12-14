package main

import "fmt"

func main() {
	var money int = 200
	var busy bool
	// var busy bool = true
	switch money {
	case 20:
		fmt.Println("点个外卖")
	case 200:
		fmt.Println("下个馆子")
		// 如果忙，不吃了，否则再吃点零食
		if busy {
			break
		} else {
			fmt.Println("再吃点零食")
		}
	case 2000:
		fmt.Println("米其林")
	default:
		fmt.Println("容我想想")
	}
	fmt.Println("end")
}
