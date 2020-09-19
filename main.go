package main

import "fmt"

func main() {
	m := map[string]int{
		"Adrien": 7,
		"Raph":   33,
		"Aurel":  27,
		"Aby":    11,
	}

	for name, age := range m {
		fmt.Printf("name=%s, age=%d\n", name, age)
	}

	i := 1
	for name := range m {
		fmt.Printf("name=%s\n", name)
		m[name] = i
		i++
	}

	fmt.Println(m)
}
