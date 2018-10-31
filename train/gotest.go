package main

import (
	"fmt"
)
//接收一个列表参数，其实参数是一个
func div1(num []int) int {
	res := 0
	for i := range num{
		res += num[i]
	}
	fmt.Println("div:",res)
	return res
}
//接收一个变长参数，参数是多个
func div2(num...int) int {
	res := 0
	for i := range num{
		res += num[i]
	}
	fmt.Println("div2:",res)
	return res
}

//第一个函数参数，接收一个列表作为参数
func appFunc(op func([]int) int,s...int) int {
	//rs1 := []int{}
	//for i:= range s{
	//	rs1 = append(rs1,i)
	//}
	//res := op(rs1)
	res := op(s)
	fmt.Println("appFunc",res)
	return res
}

//第一个函数参数，接收多个参数作为参数
func appFunc2(op func(...int) int,s...int) int {
	//可变长参数传入之后，是[a,b,c,d,e]这种形式
	//所有func接收的参数要为[]int
	res := op(s...)
	fmt.Println("appFunc2",res)
	return res
}


func main() {
	fmt.Println("hello world")
	//-------------------------------
	fmt.Println("===============================")
	div1([]int{1,2,3,4,5,6,7,8,9,10,11,12})
	appFunc(div1,1,2,3,4,5,6,7,8,9,10,11,12)
	//-------------------------------
	fmt.Println("===============================")
	div2(1,2,3,4,5,6,7,8,9)
	appFunc2(div2,1,2,3,4,5,6,7,8,9)
}
