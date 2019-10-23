package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
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

func main() {
	http.HandleFunc("/hi", sayHello) // set router
	http.HandleFunc("/bye", sayBye) // set router
    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
