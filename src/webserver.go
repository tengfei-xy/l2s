package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func webIndex(w http.ResponseWriter, r *http.Request) {
	var retval int = 200
	var retstr string
	l := logrus.WithFields(logrus.Fields{
		"IP":   r.RemoteAddr,
		"PATH": r.RequestURI,
	})

	if r.URL.Path == "/pam" {
		if r.URL.RawQuery[0:4] == "get=" {
			retstr, retval = shorturl_get(l, r.RemoteAddr, r.URL.RawQuery[4:])
		} else {
			retstr = "Invalid URL Path"
			l.Errorf(retstr)
			retval = 404
		}
	}

	w.WriteHeader(retval)
	fmt.Fprintf(w, retstr)
}
