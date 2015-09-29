package user

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radionoiseproject/rnp-server/interfaces"
	"log"
	"net/http"
)

type handler struct {
	hub      interfaces.Hub
	upgrader *websocket.Upgrader
}

func misakaUpgradeError(w http.ResponseWriter, r *http.Request, status int, reason error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "“‘%s,’ MISAKA says, acting as if she understands what that means.”",
		reason)
}

func Handler(hub interfaces.Hub) *handler {
	return &handler{
		hub:      hub,
		upgrader: &websocket.Upgrader{Error: misakaUpgradeError},
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}

	u := New(conn, h.hub)
	u.Run()
}
