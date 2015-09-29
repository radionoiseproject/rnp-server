// Copyright (c) 2015 The Radio Noise Project Members
// See COPYING for the license terms and complete list of copyright holders

package interfaces

// The Hub handles central coordination of users, putting them into rooms,
// etc.
type Hub interface {
	// Start the hub.
	Run()

	// Shutdown the hub, waiting for it to finish.
	Done()

	// Send a message to all connected users.
	Broadcast(Message)

	// Register a user in the hub.
	//
	// The hub stores a mapping allowing messages to be sent to specific
	// users by their id. A user can't join a room until they are
	// registered.
	RegisterUser(User) error

	// Deregister a user from the hub.
	UnregisterUser(User) error
}

// The User abstracts the details of the connections.
//
// It handles handshake/hello messages, connection timeouts, etc. and is
// responsible for converting internal message types to/from wire format.
type User interface {
	// Send a message to the user over the communication channel
	Send(Message)

	// Get the user's id string (format TBD... uuid?)
	Id() (string, error)

	// Terminate the user's connection.
	//
	// TODO this should probably be renamed, since I don't think it'll
	// block like the hub done connection does.
	Done()
}

// The internal message format used for communicating between users, hub, and
// rooms.
type Message interface {
}
