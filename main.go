package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Sum result= %v\n", sum)

	eventCnt := 0
	for eventCnt < 3 {
		fmt.Println("Retrieving events")
		eventCnt++
		if eventCnt == 3 {
			fmt.Printf("Got %d notifications \n", eventCnt)
		}
	}

	i := 0
	for {
		i++
		if i%2 != 0 {
			fmt.Println("odd looping")
			continue
		}
		fmt.Println("Loopinq ...")

		if i >= 10 {
			fmt.Println("Loop END")
			break
		}
	}
}
