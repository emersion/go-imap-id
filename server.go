package id

import (
	"github.com/emersion/go-imap/common"
	"github.com/emersion/go-imap/server"
)

type Handler struct {
	Command

	serverID ID
	gotID func(conn *server.Conn, id ID)
}

func (hdlr *Handler) Handle(conn *server.Conn) error {
	if hdlr.gotID != nil {
		hdlr.gotID(conn, hdlr.Command.ID)
	}

	res := &Response{hdlr.serverID}
	return conn.WriteResp(res)
}

// NewServer enables the ID extension on an IMAP server. id is the server ID, it
// can be nil. gotID is a function that will be called when a client sends its
// own ID.
func NewServer(s *server.Server, id ID, gotID func(conn *server.Conn, id ID)) {
	s.RegisterCapability(Capability, common.ConnectedState)

	s.RegisterCommand(commandName, func() server.Handler {
		return &Handler{serverID: id, gotID: gotID}
	})
}
