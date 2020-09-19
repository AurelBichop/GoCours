package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m["hello"] = 5
	m["GoodBye"] = 10
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	fmt.Printf("key hello, value=%v\n", m["hello"])
	i := m["GoodBye"]
	fmt.Println(i)

	j, ok := m["yelo"]
	fmt.Printf("j=%v, ok=%v\n", j, ok)

	//Pour tester l'égalité
	m["beatles"] = 2
	if _, ok := m["beatles"]; ok {
		fmt.Println("Beatles key exist")
	}

	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	delete(m, "beatles")
	fmt.Printf("Map content %v (len=%v)\n", m, len(m))

	m2 := m
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))

	m2["help"] = 44

	fmt.Printf("m content %v (len=%v)\n", m, len(m))
	fmt.Printf("m2 content %v (len=%v)\n", m2, len(m2))
}
