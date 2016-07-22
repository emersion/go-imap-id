package id

import (
	"github.com/emersion/go-imap/common"
)

// An ID command.
// See RFC 2971 section 3.1.
type Command struct {
	ID ID
}

func (cmd *Command) Command() *common.Command {
	return &common.Command{
		Name: commandName,
		Arguments: []interface{}{
			common.FormatParamList(map[string]string(cmd.ID)),
		},
	}
}

func (cmd *Command) Parse(fields []interface{}) (err error) {
	cmd.ID, err = parseID(fields)
	return
}
