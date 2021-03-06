package main

import (
	"fmt"
	"os"

	"training.go/CoursGo/HANGMAN/dictionary"
	"training.go/CoursGo/HANGMAN/hangman"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	g,_ := hangman.New(8, dictionary.PickWord())

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from the terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
