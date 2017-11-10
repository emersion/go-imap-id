package id

import (
	"github.com/emersion/go-imap/server"
)

type User interface {
	SetClientID(id ID)
}

type Conn interface {
	ID() ID

	setID(id ID)
}

type conn struct {
	server.Conn

	id ID
}

func (conn *conn) ID() ID {
	return conn.id
}

func (conn *conn) setID(id ID) {
	conn.id = id
}

type Handler struct {
	Command

	ext *extension
}

func (hdlr *Handler) Handle(conn server.Conn) error {
	conn.Server().ForEachConn(func(extended server.Conn) {
		if extended.Context() == conn.Context() {
			if connId, ok := extended.(Conn); ok {
				connId.setID(hdlr.Command.ID)
			}
		}
	})

	if user, ok := conn.Context().User.(User); ok {
		user.SetClientID(hdlr.Command.ID)
	}

	return conn.WriteResp(&Response{hdlr.ext.serverID})
}

type extension struct {
	serverID ID
}

func (ext *extension) Capabilities(c server.Conn) []string {
	return []string{Capability}
}

func (ext *extension) Command(name string) server.HandlerFactory {
	if name != commandName {
		return nil
	}

	return func() server.Handler {
		return &Handler{ext: ext}
	}
}

func (ext *extension) NewConn(c server.Conn) server.Conn {
	return &conn{Conn: c}
}

func NewExtension(serverID ID) server.Extension {
	return &extension{serverID}
}
