package main

import (
	"errors"
	"fmt"
)

func Error() {
	err := errors.New("Terjadi kesalahan")
	if err != nil {
		fmt.Println(err)
	}
}

type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func CustomError() {
	err := &MyError{"Operasi gagal"}
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	Error()
	CustomError()
}
