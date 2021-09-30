package http

import (
	"context"
	"net/http"
)

//HttpInterface ...
type HttpInterface interface {
	Do(ctx context.Context, request *http.Request) ([]byte, int, error)
	WithTimeout(ctx context.Context) (context.Context, context.CancelFunc)
}
