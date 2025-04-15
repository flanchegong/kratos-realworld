package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("123456")
	spew.Dump(s)
	panic(1)
}
