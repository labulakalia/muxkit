package http

import (
	"context"
	"net/http"
)

type CtxKeyType string

const CtxKeyRequest CtxKeyType = "http_request"

func Request2Context(r *http.Request) context.Context {
	ctx := context.WithValue(r.Context(), CtxKeyRequest, r)
	// TODO
	return ctx
}
