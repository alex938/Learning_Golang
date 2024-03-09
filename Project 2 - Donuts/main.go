package main

import (
	"fmt"
	"time"
)

func main() {

	donut_menu()
	todays_donut := bake_a_donut()
	fmt.Println(todays_donut)
	time.Sleep(3000 * time.Millisecond) 
	todays_donut.add_a_flavour()
}