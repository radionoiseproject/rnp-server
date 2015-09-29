// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package user

import "github.com/radionoiseproject/rnp-server/interfaces"
import "github.com/gorilla/websocket"

type Message struct {
}

type user struct {
	id   string
	name string
	ws   *websocket.Conn
	hub  *interfaces.Hub
}

func New(ws *websocket.Conn, hub *interfaces.Hub) *user {
	return &user{ws: ws, hub: hub}
}

// interfaces.User

func (u *user) Send(msg interfaces.Message) {
}

func (u *user) Id() (string, error) {
	err := nil
	if u.id == nil {
		err = Error("User id is not set")
	}
	return u.id, err
}

func (u *user) ForceDisconnect() {
}
