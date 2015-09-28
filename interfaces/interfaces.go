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
