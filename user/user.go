// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package user

import "github.com/radionoiseproject/rnp-server/interfaces"

type Message struct {
}

type user struct {
	id   string
	name string
	hub  *interfaces.Hub
}

func New(id string, name string, hub *interfaces.Hub) *user {
	return &User{id: id, name: name, hub: hub}
}
