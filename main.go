package main

import "fmt"

type Adresse struct {
	street, city string
}

type Personne struct {
	Name string
	Age  int
	Addr Adresse
}

func main() {
	var p Personne
	p.Name = "Bod"
	p.Age = 18
	p.Addr.city = "Perpignan"

	fmt.Println(p)
}
