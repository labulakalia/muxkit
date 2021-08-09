package http

import (
	"context"
	"net/http"
)

type CtxKeyType string

const (
	KeyHttpRequest CtxKeyType = "http_request"
	KeyHttpHeader
)

func Request2Context(r *http.Request) context.Context {
	ctx := context.WithValue(r.Context(), KeyHttpRequest, r)
	ctx = context.WithValue(ctx, KeyHttpHeader, r.Header)
	return ctx
}
