package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
	"html/template"
)

func sayHello(w http.ResponseWriter ,r *http.Request) {
	r.ParseForm()
	fmt.Printf(" Type of r = %T  and value = %v",r,r)
	fmt.Println("Form =",r.Form)
	fmt.Println(
		"URL Path = ",
		r.URL.Path)
	fmt.Println("URL scheme = ",r.URL.Scheme)
	fmt.Println("url long form = ", r.Form["url_long"])
	for k,v :=range r.Form{
		fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Println("Going to send a hello response to the client")
	fmt.Fprintf(w, "Well Hello There!") // send data to client side
}


func sayBye(w http.ResponseWriter ,r *http.Request) {

	fmt.Println("Going to say bye to the client")
	fmt.Fprintf(w, "Ok Byeee!") // send data to client side
}

func sayWasup(w http.ResponseWriter ,r *http.Request) {

	fmt.Println("Going to wassup bye to the client")
	fmt.Fprintf(w, "Wassssup!") // send data to client side
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if strings.Join(r.Form["username"],"") == "gajan"{
			if strings.Join(r.Form["password"],"") == "123" {
			fmt.Fprintf(w, "Welcome Master!") // send data to client side
		} else{
			fmt.Fprintf(w, "Master your password is wrong!") // send data to client side
		}
		} else {
			fmt.Fprintf(w, "Halt Fool!") // send data to client side
		}

    }
}

type myMux struct{
}

func (p *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/"{
		sayWasup(w,r)
		return
	}
	http.NotFound(w,r)
	return
}

func main() {
	//mux:=&myMux{}
	http.HandleFunc("/hi", sayHello) // set router
	http.HandleFunc("/bye", sayBye) // set router
	http.HandleFunc("/login", login) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port. default handler DefaultServeMux called since nil is passed.
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
/*	err1 :=http.ListenAndServe(":9091", mux)
	if err1 != nil {
       log.Fatal("ListenAndServe: ", err1)
	}
*/
}
