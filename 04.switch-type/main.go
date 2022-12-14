package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var money interface{} = 10
	var money interface{} = 10.0
	// var money interface{} = "10"
	switch money.(type) {
	case int:
		fmt.Println("money是int")
	case int64:
		fmt.Println("money是int64")
	case float64:
		fmt.Println("money是float64")
	case float32:
		fmt.Println("money是float32")
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println("end", money, reflect.TypeOf(money))
	money = 3
	fmt.Println("end", money, reflect.TypeOf(money))
}
