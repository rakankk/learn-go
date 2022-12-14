package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("nihao,golang")
	}

	fmt.Println("round 2")
	j := 0
	for ; j < 5; j++ {
		fmt.Println("round 2,hello,golang")
	}

	fmt.Println("round 3")
	l := 0
	for {
		fmt.Println("round 3")
		l++
		if l >= 3 {
			break
		}
	}

	fmt.Println("round4")
	m := 0
	for {
		fmt.Println("round 4，hello golang")
		m++
		if m > 10 {
			break
		}
		// 当m此时的值对2取模为0是，跳过下面的执行输出，返回从头执行函数
		if m%2 == 0 {
			continue
		}
		fmt.Println("round 4 跳过，重新执行")
	}
}
