package main

import (
	"flag"
	"fmt"
	"github.com/json-iterator/go"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type Email string
type Password string

type User struct {
	name string
	email Email
	password Password
}

func AllUsers() (map[string]User, error) {
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

func MainHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	postMethod := r.Method
	if postMethod != "post" {
		return
	}
	defer fmt.Printf("%s", "panic catched")
	switch postMethod {
	case "post", "POST":
		fmt.Fprintf(w,"<h1>Post from %s</h1>", r.URL)
	case "get", "GET":
		fmt.Fprintf(w,"<h1>Get from %s</h1>", r.URL)
	default:
		fmt.Fprintf(w,"<h1>Nor post or get request</h1>")
	}
}

func UserList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	allOfUsers, _ := AllUsers()
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

func EmailList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	allOfUsers, _ := AllUsers()
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
	var port string
	var host string

	cmdLine := flag.NewFlagSet("http_server", flag.PanicOnError)
	cmdLine.StringVar(&port, "p", "8088", "端口号，默认为8088")
	cmdLine.StringVar(&host, "h", "127.0.0.1", "主机名，默认127.0.0.1")
	cmdLine.Parse(os.Args[1:])
	//flag.StringVar(&port, "p", "8088", "端口号，默认为8088")
	//flag.StringVar(&host, "h", "127.0.0.1", "主机名，默认127.0.0.1")
	//flag.Parse()


	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "http base")
		flag.PrintDefaults()
	}

	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/users", UserList)
	http.HandleFunc("/emails", EmailList)
	addr := host + ":" + port
	log.Printf(addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}