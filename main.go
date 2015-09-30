// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/radionoiseproject/rnp-server/hub"
	"github.com/radionoiseproject/rnp-server/user"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	bind = flag.String("bind", ":8080", "address and port to listen on")
)

type misakaNotFoundHandler struct{}

func (_ misakaNotFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "“‘You seem to be a bit lost,’ MISAKA says as she looks on with a concerned expression.”")
}

func main() {
	flag.Parse()

	h := hub.New()
	go h.Run()

	router := mux.NewRouter()
	router.NotFoundHandler = misakaNotFoundHandler{}
	router.Handle("/rnp", user.Handler(h))

	loggedRouter := handlers.CombinedLoggingHandler(os.Stderr, router)

	http.Handle("/", loggedRouter)

	var l net.Listener
	var listenFields log.Fields
	if strings.ContainsRune(*bind, '/') {
		listenFields = log.Fields{"bind": *bind, "listener": "unix"}
		a, err := net.ResolveUnixAddr("unix", *bind)
		if err != nil {
			log.WithFields(listenFields).Fatal(err)
		}
		l, err = net.ListenUnix("unix", a)
		if err != nil {
			log.WithFields(listenFields).Fatal(err)
		}
	} else {
		listenFields = log.Fields{"bind": *bind, "listener": "tcp"}
		a, err := net.ResolveTCPAddr("tcp", *bind)
		if err != nil {
			log.WithFields(listenFields).Fatal(err)
		}
		l, err = net.ListenTCP("tcp", a)
		if err != nil {
			log.WithFields(listenFields).Fatal(err)
		}
	}

	log.WithFields(listenFields).Info("Radio Noise Project listening")

	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}

	h.Done()
}
