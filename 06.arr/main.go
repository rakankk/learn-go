package main

import (
	"fmt"
)

func main() {
	var age int = 35
	fmt.Println(age)
	var ages [5]int = [5]int{34, 56, 35, 46, 35}
	fmt.Println(ages)
	var ages1 = [5]int{34, 56, 35, 46, 35}
	fmt.Println(ages1)
	ages2 := [5]int{34, 56, 35, 46, 35}
	fmt.Println(ages2)

	a := [3]int{1, 2, 3}
	var b [3]int = [3]int{}
	//var c [3]int = [4]int{} //编译不通过，类型不匹配错误
	d := [5]int{}
	e := [...]int{1, 2, 3, 4} //golang会检测后面数组长度，补到前面
	fmt.Println(a)
	fmt.Println(b)
	//	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//数组赋值
	var i [3]int
	i = [3]int{1, 2, 3}
	fmt.Println(i)
	i = [3]int{2, 3, 4}
	fmt.Println(i)

	j := [3]int{}
	j[0] = 1
	j[1] = 2
	j[2] = 3
	fmt.Println(j)

	var ages4 [3]int
	fmt.Println("ages4", ages4)
	ages4[0] = 1000
	ages4[1] = 1111
	ages4[2] = 2222
	fmt.Println("ages4", ages4)

	for k := 0; k < len(ages4); k++ {
		fmt.Println(ages4[k])
	}
	// 换种range写法，上面的循环
	for k, val := range ages4 {
		fmt.Println(ages4[k], "====>", k, "->", val)
	}
}
