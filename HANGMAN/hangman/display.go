package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	/$$                                                                             /$$                        /$$$$$$  /$$$$$$$ 
	| $$                                                                            | $$                       /$$__  $$| $$__  $$
	| $$$$$$$   /$$$$$$  /$$$$$$$   /$$$$$$  /$$$$$$/$$$$   /$$$$$$  /$$$$$$$       | $$$$$$$  /$$   /$$      | $$  \ $$| $$  \ $$
	| $$__  $$ |____  $$| $$__  $$ /$$__  $$| $$_  $$_  $$ |____  $$| $$__  $$      | $$__  $$| $$  | $$      | $$$$$$$$| $$$$$$$ 
	| $$  \ $$  /$$$$$$$| $$  \ $$| $$  \ $$| $$ \ $$ \ $$  /$$$$$$$| $$  \ $$      | $$  \ $$| $$  | $$      | $$__  $$| $$__  $$
	| $$  | $$ /$$__  $$| $$  | $$| $$  | $$| $$ | $$ | $$ /$$__  $$| $$  | $$      | $$  | $$| $$  | $$      | $$  | $$| $$  \ $$
	| $$  | $$|  $$$$$$$| $$  | $$|  $$$$$$$| $$ | $$ | $$|  $$$$$$$| $$  | $$      | $$$$$$$/|  $$$$$$$      | $$  | $$| $$$$$$$/
	|__/  |__/ \_______/|__/  |__/ \____  $$|__/ |__/ |__/ \_______/|__/  |__/      |_______/  \____  $$      |__/  |__/|_______/ 
								   /$$  \ $$                                                   /$$  | $$                          
								  |  $$$$$$/                                                  |  $$$$$$/                          
								   \______/                                                    \______/                           	
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		__________
		||/		  |
		||		  |	
		||		  o
		||		 \î/
		||		 /i\	
		||
		==================
		==================
		`
	case 1:
		draw = `
		__________
		||/		  |
		||		  |	
		||		  o
		||		 
		||		 
		||
		==================
		==================
		`
	case 2:
		draw = `
		__________
		||/		  
		||		  	
		||		 
		||		 
		||		 	
		||
		==================
		==================
		`
	case 3:
		draw = `
		
		||/		 
		||		 	
		||		 
		||		 
		||		 	
		||
		==================
		==================
		`
	case 4:
		draw = `

		||		 
		||		 
		||		 	
		||
		==================
		==================
		`
	case 5:
		draw = `
	 
		||		 	
		||
		==================
		==================
		`
	case 6:
		draw = `
		==================
		==================
		`
	case 7:
		draw = `
		==================
		`
	case 8:
		draw = `
		
		
		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good Guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter %s is already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, %s is not in the word", guess)
	case "lost":
		fmt.Print("You lost the word was :")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You WON !!! the word was :")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
