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
	"os/signal"
	"strings"
	"syscall"
)

var (
	bind     = flag.String("bind", ":8080", "address and port to listen on")
	loglevel = flag.String("loglevel", "info", "level of log output (debug, info, warn, error, fatal, panic)")
)

type misakaNotFoundHandler struct{}

func (_ misakaNotFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "“‘You seem to be a bit lost,’ MISAKA says as she looks on with a concerned expression.”")
}

func main() {
	flag.Parse()

	// Logging setup
	level, err := log.ParseLevel(*loglevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	h := hub.New()
	go h.Run()

	router := mux.NewRouter()
	router.NotFoundHandler = misakaNotFoundHandler{}
	router.Handle("/rnp", user.Handler(h))

	logHandler := handlers.CombinedLoggingHandler(os.Stderr, router)
	proxyHeaderHandler := handlers.ProxyHeaders(logHandler)

	http.Handle("/", proxyHeaderHandler)

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
		defer func() {
			log.WithFields(listenFields).Debug("Cleaning up socket")
			err = os.Remove(*bind)
			if err != nil {
				log.WithFields(listenFields).Fatalf(
					"Failed to unlink socket: %s", err)
			}
		}()
		err = os.Chmod(*bind, os.ModePerm)
		if err != nil {
			log.WithFields(listenFields).Fatalf("Could not set socket permissions: %s", err)
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

	go func() {
		if err := http.Serve(l, nil); err != nil {
			log.Fatal(err)
		}
	}()
	log.WithFields(listenFields).Info("Radio Noise Project listening")

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-sigChan
	log.WithFields(log.Fields{"signal": sig}).Info("Exiting due to signal")

	h.Done()
}
