package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

type hwinfo map[string]string

var (
	isServer bool
)

// type serverConfig struct {
// cpu hwInfo
// // mem hwInfo
// // bus hwInfo
// // nic hwInfo
// // user swInfo
// }

// http service to receive configuration from cluster machines
type serverConfigService struct{}

func init() {
	flag.BoolVar(&isServer, "s", false, "Run fenir as the server")
}

/*
 * json.Decoder is preferred since request data comes from an io.Reader
 */
func serverParseAll(r *http.Request) {
	// cpuinfo := hwinfo{}
	var cpuinfo map[string]string
	dec := json.NewDecoder(r.Body)
	dec.Decode(&cpuinfo)
	fmt.Println(cpuinfo)
}

func (svc serverConfigService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverParseAll(r)
	fmt.Fprintf(w, "Hi there")
}

func main() {
	flag.Parse()

	if !isServer {
		log.Println("Client: collect server info")
		return
	}

	l, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	svc := serverConfigService{}
	http.Serve(l, svc)
}
