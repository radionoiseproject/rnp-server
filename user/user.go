// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package user

import "github.com/radionoiseproject/rnp-server/interfaces"
import "github.com/gorilla/websocket"

type Message struct {
}

type Error string

func (e Error) Error() string {
	return string(e)
}

type user struct {
	id          string
	name        string
	initialized bool
	ws          *websocket.Conn
	hub         interfaces.Hub
}

func New(ws *websocket.Conn, hub interfaces.Hub) *user {
	return &user{ws: ws, hub: hub}
}

func (u *user) Run() {
	// TODO: start the user handler goroutine thing
}

// interfaces.User

func (u *user) Send(msg interfaces.Message) {
}

func (u *user) Id() (string, error) {
	if u.initialized {
		return u.id, nil
	} else {
		return u.id, Error("User has not completed initialization")
	}
}

func (u *user) ForceDisconnect() {
}
