package main

import (
	"fmt"

	"training.go/GoCours/player"
)

func main() {
	var p1 player.Player
	p1.Name = "Adrien"
	p1.Age = 10

	// une maniere (toute les variables doivent être renseigné)
	a := player.Avatar{"http://avatar.jpg"}

	fmt.Printf("avatar = %v\n", a)

	//une autre
	p3 := player.Player{
		Name: "John",
		Age:  25,
		Avatar: player.Avatar{
			Url: "http://url.com",
		},
	}

	fmt.Printf("player3 = %v\n", p3)

	p4 := player.New("Bobette")
	p4.Avatar = a
	fmt.Printf("player4 = %v\n", p4)
}
