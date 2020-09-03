package main

import (
	"fmt"
)

func main() {
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -2

	fmt.Printf("%v,,  len=%d, cap=%d\n", nums, len(nums), cap(nums))

	nums = append(nums, 64)
	fmt.Printf("%v,,  len=%d, cap=%d\n", nums, len(nums), cap(nums))

	nums = append(nums, -8)
	fmt.Printf("%v,,  len=%d, cap=%d\n", nums, len(nums), cap(nums))

	letters := []string{"g", "o", "l", "a", "n", "g"}
	//letters = append(letters, "!")
	fmt.Printf("%v,,  len=%d, cap=%d\n", letters, len(letters), cap(letters))

	sub1 := letters[0:2]
	sub2 := letters[2:]

	fmt.Printf("%v,,  len=%d, cap=%d\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v,,  len=%d, cap=%d\n", sub2, len(sub2), cap(sub2))

}
