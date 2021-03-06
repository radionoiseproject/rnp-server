// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package hub

import "github.com/radionoiseproject/rnp-server/interfaces"
import "fmt"

type message interface {
	String() string
}

type hub struct {
	users   map[string]interfaces.User
	channel chan message
	shutdown chan struct{}
}

func New() *hub {
	return &hub{
		users:   make(map[string]interfaces.User),
		channel: make(chan message),
		shutdown: make(chan struct{}),
	}
}

// interfaces.Hub

func (h *hub) Run() {
	for m := range h.channel {
		switch m.(type) {
		case *broadcastMessageType:
			h.handleBroadcastMessage(m.(*broadcastMessageType))
		case *registerUserMessageType:
			h.handleRegisterUserMessage(m.(*registerUserMessageType))
		default:
			fmt.Printf("Unexpected hub message: %s", m)
		}
	}
	h.shutdown <- struct{}{}
}

func (h *hub) Done() {
	close(h.channel)
	<-h.shutdown
}

func (h *hub) Broadcast(msg interfaces.Message) {
	h.channel <- broadcastMessage(msg)
}

func (h *hub) RegisterUser(user interfaces.User) error {
	id, err := user.Id()
	if err != nil {
		return err
	}
	h.channel <- registerUserMessage(id, user)
	return nil
}

func (h *hub) UnregisterUser(u interfaces.User) error {
	return nil
}
