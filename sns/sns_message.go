package sns

import (
	"bytes"
	"reflect"
)

// MessageSNS implements the Message Interface
type MessageSNS struct {
	*BaseMessage
	// other properties of the sns message
}

func NewSNSMessage() *MessageSNS {
	return &MessageSNS{
		&BaseMessage{
			headers:     make(map[string]interface{}),
			headerTypes: make(map[string]reflect.Kind),
			body:        &bytes.Buffer{},
		},
	}
}

func (lm *MessageSNS) Rsvp(yes bool, options ...Option) (err error) {
	return
}
