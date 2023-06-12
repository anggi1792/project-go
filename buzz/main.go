package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(Getdata(i))
	}

}

func Getdata(i int) string{
if i%3 == 0 && i%5 == 0 {
	return "fizzbuzz"
 }else if i%3 == 0 {
 	return "fizz"
 }else if i%5 == 0 {
 	return "buzz"
 }else if i%15 == 0 {
 	return "Fizz Buzz"
 }else {
 	return fmt.Sprintf("%d", i)
 }

}
