package main

import (
	"fmt"
)

func main() {
	names := []string{"Aby", "Raph", "Adrien", "Aurel"}
	for i, n := range names {
		fmt.Printf("Username=%s (index=%d)\n", n, i)
	}

	for _, c := range "golang" {
		fmt.Printf("%v\n", string(c))
	}
}
