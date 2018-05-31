package id

import (
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/responses"
)

// An ID response.
// See RFC 2971 section 3.2.
type Response struct {
	ID ID
}

func (r *Response) Handle(resp imap.Resp) (err error) {
	name, fields, ok := imap.ParseNamedResp(resp)
	if !ok || name != responseName {
		return responses.ErrUnhandled
	}

	r.ID, err = parseID(fields)

	return
}

func (r *Response) Parse(fields []interface{}) (err error) {
	r.ID, err = parseID(fields)
	return
}

func (r *Response) WriteTo(w *imap.Writer) error {
	fields := []interface{}{responseName, formatID(r.ID)}

	res := imap.NewUntaggedResp(fields)
	return res.WriteTo(w)
}
