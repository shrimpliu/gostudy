package main

import ("fmt")

func main() {
	fmt.Println("开始学习GO语言")
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}