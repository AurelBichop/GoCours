package main

import (
	"fmt"
)

func main() {
	var day int = 25

	if day < 15 {
		fmt.Printf("We 're in the first half of thr month (day=%d)\n", day)
	} else if day == 18 {
		fmt.Printf("We are Special Day (day=%d)\n", day)
	} else {
		fmt.Printf("We are second half of the month (day=%d)\n", day)
	}

	year, month, day := 20010, 110, 100

	fmt.Printf("Date=%d/%d/%d\n", day, month, year)

	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("The BirthDay for the Golang")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("A peut prêt la date")
	} else {
		fmt.Println("Just a another Day")
	}
}
