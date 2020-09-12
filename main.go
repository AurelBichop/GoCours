package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%v - %v", p.Title, p.Text[:50])
}

func (p *Post) Published() bool {
	return p.published
}

func (p *Post) Publish() {
	p.published = true
}

func (p *Post) UnPublished() {
	p.published = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {
	p := Post{
		Title: "Go releaser",
		Text: `Lorem ipsum edsz fzefze f ze
efze fez fzef zef ze f f zez  ezf ezfzef ezf zfpezkfopzejoàfuziçfh uguzgfygzeyuqafgqeyfgqzgg qfg uiq`,
	}

	fmt.Println(p.Headline())
	fmt.Printf("Post Published %v\n", p.Published())

	p.Publish()
	fmt.Printf("Post Published %v\n", p.Published())

	pythonPost := &Post{
		Title: "Python releaser",
		Text: `Lorem ipsum edsz fzefze f ze
efze fez fzef zef fuhdhsfgyfd ezf ezfzef ezf zfpezkfopzejoàfuziçfh uguzgfygzeyuqafgqeyfgqzgg qfg uiq`,
	}
	UpperTitle(pythonPost)
	fmt.Println(pythonPost.Headline())
}
