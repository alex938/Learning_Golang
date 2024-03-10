package main

import "fmt"

func main() {
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}
	for _, num := range numbers {
		if _is_even(num) {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}

func _is_even(number int) bool {
	//if  evaluation of number%2 == 0, return True else False
	return number%2 == 0
}