package main

import (
	"fmt"
	"log"
	"net/http"
	"errors"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
)

type Email string
type Password string

type User struct {
	name string
	email Email
	password Password
}

func mainHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := ps.ByName("user")
	userInfo, err := searchUser(userName)
	if err != nil {
		//errInfo := map[string]interface{}{"message": err.Error()}
		//var json = jsoniter.ConfigCompatibleWithStandardLibrary
		//user, encodeErr := json.Marshal(&errInfo)
		//if encodeErr != nil {
			fmt.Fprintf(w, "%s", err)
		//}
		//w.Write(user)
	}
	//var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//user, encodeErr := json.Marshal(&userInfo)
	//if encodeErr != nil {
	//	fmt.Errorf("%s", encodeErr)
	//}
	fmt.Fprintf(w, "%s", userInfo)
	//w.Write(user)
	w.Write()
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

func searchUser(userName string) (User, error) {
	userList, err := allUsers()
	if err != nil {
		return User{}, err
	}
	for name, info := range userList {
		if name == userName {
			return info,nil
		}
	}
	return User{}, errors.New("User not exist")
	}

func userList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func emailList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	router := httprouter.New()
	router.GET("/", mainHandler)
	router.GET("/home/:user", home)
	router.GET("/users", userList)
	router.GET("/emails", emailList)
	err := http.ListenAndServe(":8088", router)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}


