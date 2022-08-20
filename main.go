package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter no.")
	fmt.Scanln(&num)
	fmt.Println(evenodd(num))
}

func evenodd(num int) string {
	var result string
	if num%2 == 0 {
		result = "even"
	} else {
		result = "odd"
	}
	return result
}
