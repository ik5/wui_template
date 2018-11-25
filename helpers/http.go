package helpers

import "net/http"

// HTTPHandler holds the callback structure for the handler
type HTTPHandler func(handle http.Handler) http.Handler
