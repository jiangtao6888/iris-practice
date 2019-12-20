package Exceptions

import (
	"fmt"
	"iris/app/Code"
)

type Exception struct {
	code    int64
	message string
}

func (e *Exception) GetCode() int64 {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) String() string {
	return fmt.Sprintf("(%d) %s", e.code, e.message)
}

func Desc(code int64) string {
	if e, ok := Code.Message[code]; ok {
		return e
	}
	return "server internal error"
}
