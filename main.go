package main

/*
	Atlas Logging Service
	Thijs Haker
*/

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Environment/lib"
)

var msgBuf []string

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas Logging Service %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(*netAddr, nil)
}
