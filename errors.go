package goninjam

import "fmt"

type AuthorizationFailed struct {
	User string
}

func (e *AuthorizationFailed) Error() string {
	return fmt.Sprintf("Authoriazation failed for user '%s'", e.User)
}
