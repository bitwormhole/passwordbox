package errors

import (
	"errors"
	"fmt"

	"github.com/bitwormhole/passwordbox/core/loggers"
)

func HandleError(err error) {
	if err != nil {
		return
	}
	loggers.LogE("errors.handle(err) : %s", err.Error())
}

func HandlePanic(x any) {
	if x == nil {
		return
	}
	err := PanicToError(x)
	if err == nil {
		return
	}
	loggers.LogE("errors.handle(panic) : %s", err.Error())
}

func PanicToError(x any) error {

	if x == nil {
		return nil
	}

	// error
	err, ok := x.(error)
	if ok {
		return err
	}

	// string
	str, ok := x.(string)
	if ok {
		return errors.New(str)
	}

	// other
	return fmt.Errorf("panic:%v", x)
}
