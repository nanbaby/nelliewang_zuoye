package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 输入输出函数
// add by nelliewang
func InputNum()(checkNum1, checkNum2 int, err error){
	fmt.Println("Input the first number(Fizz):")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	temp := input.Text()
	checkNum1, err = strconv.Atoi(temp)
	if err != nil {
		return
	}
	fmt.Println("Input the second number(Buzz):")
	input.Scan()
	temp = input.Text()
	checkNum2, err = strconv.Atoi(temp)
	return
}

func OutputNum(outNum []string)(err error){
	fmt.Println("第*个人    报数 ")
	for i:=1; i<len(outNum); i++ {
		fmt.Println(i, outNum[i])
	}
	return
}