package main

import (
	"fmt"
)

func UpdateVal(val string) {
	val = "value"
}

func UpdatePtr(ptr *string) {
	*ptr = "pointer"
}

func main() {

	i := 69
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("*p=%v\n", *p)
	fmt.Println("------------------")

	s := "YOLO"
	sPtr := &s
	var s2 string = *sPtr

	fmt.Println("String Pointer")

	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)

	fmt.Println("------------------")

	*sPtr = "Yann"
	fmt.Println("Dereference and Update")

	fmt.Printf("s=%v\n", s)
	fmt.Printf("sPtr=%v\n", sPtr)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Printf("s2=%v\n", s2)
	fmt.Println("------------------")

	UpdateVal(s)
	fmt.Println("Func Update Val")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("------------------")

	//UpdatePtr(sPtr)
	UpdatePtr(&s)
	fmt.Println("Func Update Pointer")
	fmt.Printf("s=%v\n", s)
	fmt.Printf("*sPtr=%v\n", *sPtr)
	fmt.Println("------------------")
}
