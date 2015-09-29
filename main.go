// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/radionoiseproject/rnp-server/hub"
	"github.com/radionoiseproject/rnp-server/user"
	"log"
	"net/http"
	"os"
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

	log.Printf("Radio Noise Project listening on %s\n", *bind)

	if err := http.ListenAndServe(*bind, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	h.Done()
}
