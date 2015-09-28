// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package hub

import "github.com/radionoiseproject/rnp-server/interfaces"

type broadcastMessageType struct {
	msg interfaces.Message
}

func (m *broadcastMessageType) String() string {
	return "broadcastMessage"
}

func broadcastMessage(msg interfaces.Message) message {
	return &broadcastMessageType{msg: msg}
}

func (h *hub) handleBroadcastMessage(m *broadcastMessageType) {
	for _, user := range h.users {
		user.Send(m.msg)
	}
}
