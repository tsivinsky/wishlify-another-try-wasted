package router

import "net/http"

type Handler = func(ctx context) http.HandlerFunc
