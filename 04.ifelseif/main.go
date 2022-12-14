package main

import "fmt"

func main() {
	var rmb int = 21
	if rmb <= 20 {
		fmt.Println("点个外卖")
	} else if rmb <= 200 {
		// 如果有不超过20块钱，点个外卖
		fmt.Println("下个馆子")
	} else if rmb <= 2000 {
		// 如果不超过2000，去新马泰
		fmt.Println("去米其林")
	} else if rmb <= 20000 {
		fmt.Println("去新马泰")

	} else {
		fmt.Println("容我想想")
	}
	// 如果不超过200，下个馆子
	// 如果不超过2000，去米其林
	// 如果不超过20000，去新马泰
	// 如果再多，容我想想
	// 在整个函数里面，程序从上到下依次执行，满足条件后就会执行，后面执行后跳出if直接end
}
