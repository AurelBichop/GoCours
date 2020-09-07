package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Processline search for old in line to replace it by new
// It returns found=true, if the pattern was found, res with the resulting string
// and occ with the number of the occurence of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line

	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}

// FindReplaceFile
func FindReplaceFile(src, dst string, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)

	defer writer.Flush()

	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}

		fmt.Fprintf(writer, res)
		lineIdx++
	}

	return occ, lines, nil
}

func main() {
	old := "Baker"
	new := "Baker-Bichotte"

	occ, lines, err := FindReplaceFile("text.txt", "texte.txt", old, new)
	if err != nil {
		fmt.Printf("Erro while executing find replace : %v\n", err)
	}

	fmt.Println("== Summary ==")
	defer fmt.Println("== End of Summary ==")

	fmt.Printf("Number of occurences of %v: %v\n", old, occ)
	fmt.Printf("Number of lines : %d\n", len(lines))

	//Affichage pour lines
	fmt.Print("Lines [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}

	}
	fmt.Println(" ]")
	//=======================

}
