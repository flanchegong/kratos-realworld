package server

import (
	nethttp "net/http"

	"github.com/go-kratos/kratos-layout/internal/errors"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
