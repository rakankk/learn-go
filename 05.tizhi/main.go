package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("姓名")
	fmt.Scanln(&name)

	var weight float64 = 100.0
	fmt.Print("体重（kg）")
	fmt.Scanln(&weight) // &weight表示取这个盒子的编号，去取值，录入

	var tall float64 = 1.76
	fmt.Print("身高（m）")
	fmt.Scanln(&tall)

	var bmi float64 = weight / (tall * tall)
	var age int = 26
	fmt.Print("年龄（岁）")
	fmt.Scanln(&age)

	var sexWight int
	var sex string = "男"
	fmt.Print("性别（男|女）")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWight = 1
	} else {
		sexWight = 0
	}
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWight)) / 100
	fmt.Println("体脂率是：", fatRate)

	if sex == "男" {
		// 编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("标准")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println("偏胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}

		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				fmt.Println("标准")
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				fmt.Println("偏胖")
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}
		} else if age > 60 {
			if fatRate <= 0.13 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				fmt.Println("标准")
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				fmt.Println("偏胖")
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}
		} else {
			fmt.Println("小于18岁，体脂率变化太大")
		}
	} else {
		// 编写女性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				fmt.Println("标准")
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				fmt.Println("偏胖")
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				fmt.Println("标准")
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				fmt.Println("偏胖")
			} else if fatRate > 0.35 && fatRate <= 0.40 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}
		} else if age > 60 {
			if fatRate <= 0.22 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				fmt.Println("标准")
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				fmt.Println("偏胖")
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("即刻开始健身")
			}
		} else {
			fmt.Println("小于18岁，体脂率变化太大")
		}
	}
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）?")
	fmt.Scanln(&whetherContinue)
	if whetherContinue == "y" {
		var name string
		fmt.Print("姓名")
		fmt.Scanln(&name)

		var weight float64 = 100.0
		fmt.Print("体重（kg）")
		fmt.Scanln(&weight) // &weight表示取这个盒子的编号，去取值，录入

		var tall float64 = 1.76
		fmt.Print("身高（m）")
		fmt.Scanln(&tall)

		var bmi float64 = weight / (tall * tall)
		var age int = 26
		fmt.Print("年龄（岁）")
		fmt.Scanln(&age)

		var sexWight int
		var sex string = "男"
		fmt.Print("性别（男|女）")
		fmt.Scanln(&sex)

		if sex == "男" {
			sexWight = 1
		} else {
			sexWight = 0
		}
		var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWight)) / 100
		fmt.Println("体脂率是：", fatRate)

		if sex == "男" {
			// 编写男性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.1 && fatRate <= 0.16 {
					fmt.Println("标准")
				} else if fatRate > 0.16 && fatRate <= 0.21 {
					fmt.Println("偏胖")
				} else if fatRate > 0.21 && fatRate <= 0.26 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}

			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.11 && fatRate <= 0.17 {
					fmt.Println("标准")
				} else if fatRate > 0.17 && fatRate <= 0.22 {
					fmt.Println("偏胖")
				} else if fatRate > 0.22 && fatRate <= 0.27 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}
			} else if age > 60 {
				if fatRate <= 0.13 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.13 && fatRate <= 0.19 {
					fmt.Println("标准")
				} else if fatRate > 0.19 && fatRate <= 0.24 {
					fmt.Println("偏胖")
				} else if fatRate > 0.24 && fatRate <= 0.29 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}
			} else {
				fmt.Println("小于18岁，体脂率变化太大")
			}
		} else {
			// 编写女性的体脂率与体脂状态表
			if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.2 && fatRate <= 0.27 {
					fmt.Println("标准")
				} else if fatRate > 0.27 && fatRate <= 0.34 {
					fmt.Println("偏胖")
				} else if fatRate > 0.34 && fatRate <= 0.39 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.21 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.21 && fatRate <= 0.28 {
					fmt.Println("标准")
				} else if fatRate > 0.28 && fatRate <= 0.35 {
					fmt.Println("偏胖")
				} else if fatRate > 0.35 && fatRate <= 0.40 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}
			} else if age > 60 {
				if fatRate <= 0.22 {
					fmt.Println("偏瘦")
				} else if fatRate > 0.22 && fatRate <= 0.29 {
					fmt.Println("标准")
				} else if fatRate > 0.29 && fatRate <= 0.36 {
					fmt.Println("偏胖")
				} else if fatRate > 0.36 && fatRate <= 0.41 {
					fmt.Println("肥胖")
				} else {
					fmt.Println("即刻开始健身")
				}
			} else {
				fmt.Println("小于18岁，体脂率变化太大")
			}
		}
	} else {
		return
	}

}

// go build
