package server

import (
	"encoding/json"
	"fmt"
	"testing"

	e "github.com/go-kratos/kratos-layout/internal/errors"

	"github.com/stretchr/testify/assert"
)

func TestHttpStruct(t *testing.T) {
	a := &e.HttpError{
		Errors: make(map[string][]string),
	}
	a.Errors["body"] = []string{"cannot be empty"}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s\n", string(b))
	//panic("flanche")
}
