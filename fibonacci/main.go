package main

import "fmt"


func main(){
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
}

func fibonacci(n int) (result int){
	if n < 2 {
		return n
	}

	result = fibonacci(n-1) + fibonacci(n-2)

	return
}