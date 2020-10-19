package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Adress struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Adress   Adress `json:"adress"`
}

var users = []User{
	{
		Name:     "Aurel",
		Password: "secret",
		Email:    "aurelbichop@golang.net",
		Adress: Adress{
			Street:  "rue de l'avenir",
			City:    "Purpel Rain",
			Country: "France",
		},
	},
	{
		Name:     "Raph",
		Password: "supersecret",
		Email:    "raph@golang.net",
		Adress: Adress{
			Street:  "69 rue Hadele",
			City:    "Tourcoin",
			Country: "France",
		},
	},
	{
		Name:     "Adrien",
		Password: "bigsecret",
		Email:    "adrien@golang.net",
		Adress: Adress{
			Street: "20 rue du printemps",
			City:   "Pericoin",
		},
	},
}

type PassworJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	var p PassworJsonBody
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if p.UserIndex < 0 || p.UserIndex > len(users)-1 {
		msg := fmt.Sprintf("Invalid index. got user_index=%v. valid range=[0,%v]", p.UserIndex, len(users)-1)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	u := users[p.UserIndex]

	if u.Password != p.OldPassword {
		http.Error(w, "Old password do not match", http.StatusBadRequest)
		return
	}

	if p.NewPassword != p.NewPasswordRepeat {
		http.Error(w, "password not match", http.StatusBadRequest)
		return
	}

	u.Password = p.NewPassword

	fmt.Fprintf(w, "Password update %v", u)
}

func usersList(w http.ResponseWriter, r *http.Request) {

	//Encodage du JSON
	b, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)

	}
	//Penser a sete le header
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/users", usersList)
	http.HandleFunc("/update_password", updatePassword)

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
