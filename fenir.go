package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

// http service to receive configuration from cluster machines
type serverConfigService struct{}

func (svc serverConfigService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Hi there")
	fmt.Println(string(content))
}

func main() {
	l, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// mux := http.NewServeMux()

	svc := serverConfigService{}

	http.Serve(l, svc)
}
