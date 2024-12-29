package errorsx

import (
	"fmt"
	"strings"
)

var errorFormat = `cortex error: %+v
%s`

type CortexError struct {
	message []string
	err     error
}

func (e *CortexError) Error() string {
	detail := wrapMessage(e.message...)
	return fmt.Sprintf(errorFormat, e.err, detail)
}

func Wrap(err error, message ...string) error {
	e, ok := err.(*CortexError)
	if ok {
		return e
	}

	return &CortexError{
		message: message,
		err:     err,
	}
}

func wrapMessage(message ...string) string {
	if len(message) == 0 {
		return ""
	}
	return fmt.Sprintf(`message: %s`, strings.Join(message, "\n"))
}
