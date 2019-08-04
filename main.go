package main

import (
	"fmt"
	"nelliewang_zuoye/service"
)

func main(){
	check1, check2, err := service.InputNum()
	if err != nil {
		fmt.Println("Input number error. err: ", err)
		return
	}
	outStrings := service.GetOutstrings(check1, check2)
	err = service.OutputNum(outStrings)
	if err != nil {
		fmt.Println("Output number error. err: ", err)
		return
	}
}
