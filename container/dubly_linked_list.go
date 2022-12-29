package container

import (
	"errors"
	"fmt"
)

type ValidationError struct{}

func (e *ValidationError) Error() string {
	return "validation error"
}

func huga() error {
	err := errors.New("huga")
	return err
}

func miso() error {
	return nil
}

func hoge() {
	if err := miso(); err != nil {
		var vErr *ValidationError
		if ok := errors.As(err, &vErr); ok {
			fmt.Printf("err: %v\n", vErr)
		}

		switch err.(type) {
		case interface{ Unwrap() error }:
		case interface{}:
		}
	}
}
