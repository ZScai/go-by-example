package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, C!")
	log.Printf("%s %s %s %s %s", r.RemoteAddr, r.Method, r.Proto, r.Host, r.UserAgent())
}
func main() {
	const port = 9090
	http.HandleFunc("/", sayHello)
	log.Printf("Starting HTTP server at port: %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("http server failed: ", err)
		return
	}
}
