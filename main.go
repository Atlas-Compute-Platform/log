package main

/*
	Atlas Logs
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
	fmt.Fprintf(os.Stderr, "Atlas Logs %s\n", lib.VERS)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.Usage = usage
	flag.Parse()

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/print", apiPrint)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(*netAddr, nil)
}
