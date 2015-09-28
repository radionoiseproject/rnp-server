// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package main

import "fmt"
import "github.com/radionoiseproject/rnp-server/hub"

func main() {
	h := hub.New()
	go h.Run()
	fmt.Println("Hello, World!")
}
