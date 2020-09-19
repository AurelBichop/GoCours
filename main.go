package main

import "fmt"

type User struct {
	name string
}

type Key struct {
	ID   int
	Name string
}

func main() {

	m := map[string]*User{
		"HR":  {"Bob"},
		"CEO": {"Alice"},
	}
	fmt.Println(m["HR"])

	hr := m["HR"]
	hr.name = "John"
	fmt.Println(m["HR"])

	m["CFO"] = &User{"Boubou"}
	fmt.Println(m["CFO"])

	res := make(map[Key]string)
	res[Key{1, "aze"}] = "file1"

	k := Key{2, "ert"}
	res[k] = "file2"

	delete(res, Key{1, "aze"})
	fmt.Println(res)
}
