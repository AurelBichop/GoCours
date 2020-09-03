package main

import "fmt"

func start() {
	fmt.Println("Start")
}

func stop() {
	fmt.Println("Stop")
}

func main() {
	start()
	defer stop() // LIFO last in first out

	names := []string{"Adrien", "Raph", "Aby", "Aurel"}

	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
		defer fmt.Printf("Goodbye %s\n", n)
	}
}
