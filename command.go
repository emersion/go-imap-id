package id

import (
	"github.com/emersion/go-imap"
)

// An ID command.
// See RFC 2971 section 3.1.
type Command struct {
	ID ID
}

func (cmd *Command) Command() *imap.Command {
	return &imap.Command{
		Name:      commandName,
		Arguments: []interface{}{formatID(cmd.ID)},
	}
}

func (cmd *Command) Parse(fields []interface{}) (err error) {
	cmd.ID, err = parseID(fields)
	return
}
