package main

import (
	"fmt"
)

type Rect struct {
	Width, Height int
}

//receiver function
func (r Rect) Area() int {
	return r.Width * r.Height
}

func (r Rect) Doublesize() {
	r.Width *= 2
	r.Height *= 2
	fmt.Println("In DoubleSize()", r)
}

//Pour un Affichage personnalisé
func (r Rect) String() string {
	return fmt.Sprintf("Rectangle ==> width=%v, height=%v", r.Width, r.Height)
}

func main() {
	r := Rect{2, 4}
	fmt.Printf("Restangle Aire = %v\n", r.Area())
	fmt.Println(r)

	r.Doublesize()
	fmt.Println("Main", r)
}
