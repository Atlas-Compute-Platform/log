package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiReset(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	logBuf = make([]string, 0, 160)
}

func apiList(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	for _, v := range logBuf {
		fmt.Fprintf(w, "%s\n", v)
	}
}

func apiLog(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		err error
		svc string
		msg string
		req = make(lib.Dict)
	)

	if req, err = lib.ReceiveDict(r.Body); err != nil {
		lib.LogError(os.Stderr, "main.apiLog", err)
		return
	}

	if logBuf[len(logBuf)-1] != "" {
		logBuf = make([]string, logSize)
	}

	svc = req[lib.KEY_SVC]
	msg = req[lib.KEY_MSG]
	logBuf = append(logBuf, fmt.Sprintf("%s: %s", svc, msg))
}
