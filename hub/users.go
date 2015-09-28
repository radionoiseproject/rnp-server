// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package hub

import "github.com/radionoiseproject/rnp-server/interfaces"
import "fmt"
import "log"

type registerUserMessageType struct {
	id   string
	user interfaces.User
}

func (m *registerUserMessageType) String() string {
	return fmt.Sprintf("registerUserMessage: %s", m.id)
}

func registerUserMessage(id string, user interfaces.User) message {
	return &registerUserMessageType{id: id, user: user}
}

func (h *hub) handleRegisterUserMessage(m *registerUserMessageType) {
	user, ok := h.users[m.id]
	if ok {
		log.Printf("Attempting to register already-registered user: %s",
			m.id)
		user.ForceDisconnect()
	}
	h.users[m.id] = m.user
}
