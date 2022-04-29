package main

/*
	Atlas Logging Service
	Thijs Haker
*/

import (
	"flag"
	"net/http"

	"github.com/Atlas-Compute-Platform/lib"
)

var logSize uint
var logBuf []string

func main() {
	lib.SvcName = "Atlas Logging Service"
	lib.SvcVers = "1.0"

	var netAddr = flag.String("p", lib.PORT, "Specify port")
	flag.UintVar(&logSize, "s", 2048, "Specify size")
	flag.Usage = lib.Usage
	flag.Parse()

	logBuf = make([]string, logSize)

	http.HandleFunc("/ping", lib.ApiPing)
	http.HandleFunc("/reset", apiReset)
	http.HandleFunc("/list", apiList)
	http.HandleFunc("/log", apiLog)

	http.ListenAndServe(*netAddr, nil)
}
