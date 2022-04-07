package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Environment/lib"
)

func apiReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msgBuf = make([]string, 0, 160)
}

func apiPrint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	for _, v := range msgBuf {
		fmt.Fprintf(w, "%s\n", v)
	}
}

func apiLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var (
		err error
		svc string
		msg string
		req = make(lib.Dict)
	)

	if req, err = lib.LoadDict(r.Body); err != nil {
		lib.LogError(os.Stderr, "main.apiLog", err)
		return
	}

	svc = req[lib.KV_SVC]
	msg = req[lib.KV_MSG]
	msgBuf = append(msgBuf, fmt.Sprintf("%s: %s", svc, msg))
}
