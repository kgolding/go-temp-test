package scope

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

var ErrTooShort = errors.New("too short")
var ErrInvalid = errors.New("invalid message")

type Message struct {
	Type    PagerType
	Capcode int
	Beep    PagerBeep
	Message string
}

func (m Message) String() string {
	return fmt.Sprintf("Type: %s, Capcode: %d, Beep: %s, Message: '%s'",
		m.Type.String(), m.Capcode, m.Beep.String(), m.Message)
}

func Decode(b []byte) (m Message, err error) {
	if len(b) < 11 {
		err = ErrTooShort
		return
	}

	crIndex := bytes.IndexByte(b, 0x0d) // \r
	if crIndex < 11 {
		err = ErrInvalid
		return
	}

	m.Type = PagerType(b[0])

	var u64 uint64
	u64, err = strconv.ParseUint(string(b[1:8]), 10, 32)
	if err != nil {
		return
	}
	m.Capcode = int(u64)

	m.Beep = PagerBeep(b[8])
	m.Message = string(b[9:crIndex])

	return
}
