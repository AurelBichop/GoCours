package main

import (
	"fmt"
)

func main() {
	var names [3]string
	fmt.Printf("names=%v (len=%v) \n", names, len(names))

	names[0] = "Bob"
	names[2] = "Adrien"
	fmt.Printf("names=%v (len=%v) \n", names, len(names))
	fmt.Printf("names=%v\n", names[2])

	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("odds=%v len=%d\n", odds, len(odds))
}
