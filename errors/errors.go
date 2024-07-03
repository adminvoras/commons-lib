package errors

import (
	"errors"
	"fmt"
)

func New(err error, msg string) error {
	if err != nil {
		return errors.New(fmt.Sprintf("%s:%s", msg, err.Error()))
	}
	return errors.New(msg)
}
