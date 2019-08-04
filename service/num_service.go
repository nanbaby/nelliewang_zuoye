package service

import "fmt"

const TOTLE = 100

func GetOutstrings(check1, check2 int) (ret []string) {
	ret = make([]string, TOTLE + 1)
	for i:=1; i<=TOTLE; i++ {
		out1 := getNumOut(i, check1)
		out2 := getNumOut(i,check2)
		if out1 && out2 {
			ret[i] = "FizzBuzz"
		} else if out1 {
			ret[i] = "Fizz"
		} else if out2 {
			ret[i] = "Buzz"
		} else {
			ret[i] = fmt.Sprint(i)
		}
	}
	return
}

// 获取当前数字num对于标志数字check的外显状态
// 返回ret true:显示对应的字符串， false:显示原数字
// add by nelliewang
func getNumOut(num, check int) (ret bool) {
	if num % check == 0 {
		return true
	}
	for num > 0 {
		if num % 10 == check {
			return true
		}
		num /= 10
	}
	return false
}