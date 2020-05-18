package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
)

type Email string
type Password string

type User struct {
	name string
	email Email
	password Password
}

func allUsers() (map[string]User, error) {
	var users = make(map[string]User)
	sunny := &User{
		name: "sunnywalden",
		email: "sunnywalden@gmail.com",
	}
	henry := &User{
		name: "Henry Zhang",
		email: "henryzhang@gmail.com",
	}

	fmt.Print(*sunny, *henry, "\n")
	users["sunny"] = *sunny
	users["henry"] = *henry
	return users, nil
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postMethord := r.Method
	switch postMethord {
	case "post", "POST":
		fmt.Fprintf(w,"<h1>Post from %s</h1>", r.URL)
	case "get", "GET":
		fmt.Fprintf(w,"<h1>Get from %s</h1>", r.URL)
	default:
		fmt.Fprintf(w,"<h1>Nor post or get request</h1>")
	}
}

func userList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	allOfUsers, _ := allUsers()
	fmt.Print(allOfUsers,"\n")

	var names []string
	for _, user := range allOfUsers {
		names = append(names, user.name)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonNames, encodeErr := json.Marshal(&names)
	if encodeErr != nil {
		fmt.Errorf("%s", encodeErr)
	}
	fmt.Fprintf(w, "%s", jsonNames)
}

func emailList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	allOfUsers, _ := allUsers()
	fmt.Print(allOfUsers,"\n")

	var emails []Email
	for _, user := range allOfUsers {
		emails = append(emails, user.email)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonEmails, encodeErr := json.Marshal(&emails)
	if encodeErr != nil {
		fmt.Errorf("%s", encodeErr)
	}
	fmt.Fprintf(w, "%s", jsonEmails)
}


func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/users", userList)
	http.HandleFunc("/emails", emailList)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}