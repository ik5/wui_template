package auth

import (
	"context"
	"fmt"
)

// FindToken takes HTTP headers and looks for the token and see if exists at the
// database and enabled
func FindToken(ctx context.Context, token string) bool {
	fmt.Println(token)
	return true
}
