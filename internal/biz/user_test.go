package biz

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("123456")
	spew.Dump(s)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$ZbX9Xq9e4X826mxMgeTVN..AWcF/U3GuE6ys./6ruYuz9PAmZe6jG", "1234567"))
}
