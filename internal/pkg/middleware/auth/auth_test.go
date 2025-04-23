package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerate(t *testing.T) {
	token := GenerateToken("eric")
	spew.Dump(token)
	panic(1)
	if token == "" {
		t.Error("Expected a non-empty token")
	}
}
