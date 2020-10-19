package main

import (
	"flag"
	"fmt"
	"os"

	"training.go/Dictionary/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")

	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknow action: %v\n", *action)
	}

}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("'%v' added to the dictionnary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("%v Deleted\n", word)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
