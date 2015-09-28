package hub

import "github.com/radionoiseproject/rnp-server/interfaces"
import "fmt"

type message interface {
	String() string
}

type hub struct {
	users   map[string]interfaces.User
	channel chan message
}

func New() *hub {
	return &hub{
		users:   make(map[string]interfaces.User),
		channel: make(chan message, 100),
	}
}

func (h *hub) Run() {
	for {
		m := <-h.channel
		switch m.(type) {
		case *broadcastMessageType:
			h.handleBroadcastMessage(m.(*broadcastMessageType))
		case *registerUserMessageType:
			h.handleRegisterUserMessage(m.(*registerUserMessageType))
		default:
			fmt.Printf("Unexpected hub message: %s", m)
		}
	}
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
