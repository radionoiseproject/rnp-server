// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package interfaces

type Hub interface {
	Run()
	Broadcast(Message)
	RegisterUser(User) error
	UnregisterUser(User) error
}

type User interface {
	Send(Message)
	Id() (string, error)
	ForceDisconnect()
}

type Message interface {
}
