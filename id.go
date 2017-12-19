// Implements the IMAP ID Extension, defined in RFC 2971.
package id

import (
	"errors"

	"github.com/emersion/go-imap"
)

// The ID capability.
const Capability = "ID"

const (
	commandName  = Capability
	responseName = Capability
)

// Fields identifying a client or a server. Keys MUST NOT be longer than 30
// bytes, values MUST NOT be longer than 1024 bytes. The number of fields MUST
// NOT be greater than 30.
type ID map[string]string

// Fields used in Id.
const (
	// Name of the program.
	FieldName = "name"
	// Version number of the program
	FieldVersion = "version"
	// Name of the operating system.
	FieldOS = "os"
	// Version of the operating system.
	FieldOSVersion = "os-version"
	// Vendor of the client/server.
	FieldVendor = "vendor"
	// URL to contact for support.
	FieldSupportURL = "support-url"
	// Postal address of contact/vendor.
	FieldAddress = "address"
	// Date program was released, specified as a date-time in IMAP4rev1.
	FieldDate = "date"
	// Command used to start the program.
	FieldCommand = "command"
	// Arguments supplied on the command line, if any.
	FieldArguments = "arguments"
	// Description of environment, i.e., UNIX environment variables or Windows
	// registry settings.
	FieldEnvironment = "environment"
)

func parseID(fields []interface{}) (id ID, err error) {
	if len(fields) < 1 {
		err = errors.New("No enough arguments")
		return
	}

	if fields[0] == nil {
		return
	}

	list, ok := fields[0].([]interface{})
	if !ok {
		err = errors.New("Invalid ID fields")
		return
	}

	var params map[string]string
	if params, err = imap.ParseParamList(list); err == nil {
		id = ID(params)
	}
	return
}

func formatID(id ID) interface{} {
	if id == nil {
		return nil
	}
	return imap.FormatParamListQuoted(map[string]string(id))
}
