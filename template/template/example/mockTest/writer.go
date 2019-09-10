package mockTest

import (
	"fmt"
)

type userReader interface {
	ReadUserInfo(int) int
}

type userWriter interface {
	WriteUserInfo(int)
}

type UserRepository struct {
	userReader
	userWriter
}

type realRW struct{}

func (db *realRW) ReadUserInfo(i int) int {
	return i
}

func (db *realRW) WriteUserInfo(i int) {
	fmt.Printf("put %d to db.\n", i)
}
