package hangman

import "testing"

func TestLetterInWorld(t *testing.T) {
	word := []string{"a", "d", "r", "i", "e", "n"}
	guess := "a"
	hasLetter := lettersInWord(guess, word)
	if !hasLetter {
		t.Errorf("Word %s contains letter %s. got=%v", word, guess, hasLetter)
	}
}
func TestLetterNotInWorld(t *testing.T) {
	word := []string{"a", "d", "r", "i", "e", "n"}
	guess := "b"
	hasLetter := lettersInWord(guess, word)
	if hasLetter {
		t.Errorf("Word %s does not contain letter %s. got=%v", word, guess, hasLetter)
	}
}

func TestInvalidWorld(t *testing.T) {
	_, err := New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when when using an invalid word=''")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("b")
	validState(t, "goodGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := New(3, "bob")
	//g.UsedLetters = append(g.UsedLetters, "B")
	g.MakeAGuess("b")
	g.MakeAGuess("b")

	validState(t, "alreadyGuessed", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := New(3, "bob")
	g.MakeAGuess("a")

	validState(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := New(3, "bob")
	//g.FoundLetters = []string{"B", "_", "B"}
	//g.MakeAGuess("o")
	g.MakeAGuess("b")
	g.MakeAGuess("o")
	g.MakeAGuess("b")

	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := New(3, "bob")
	//g.TurnsLeft = 1
	g.MakeAGuess("a")
	g.MakeAGuess("z")
	g.MakeAGuess("e")

	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("state should be '%v' got='%v'", expectedState, actualState)
		return false
	}
	return true
}
