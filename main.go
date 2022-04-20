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

	"github.com/Atlas-Compute-Platform/lib"
)

var logSize uint
var logBuf []string

func usage() {
	fmt.Fprintf(os.Stderr, "Atlas Logging Service %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.UintVar(&logSize, "s", 2048, "Specify size")
	flag.Usage = usage
	flag.Parse()

	logBuf = make([]string, logSize)

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(*netAddr, nil)
}
