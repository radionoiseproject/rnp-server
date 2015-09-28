package main

import "fmt"
import "github.com/radionoiseproject/rnp-server/hub"

func main() {
	h := hub.New()
	go h.Run()
	fmt.Println("Hello, World!")
}
