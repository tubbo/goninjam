package goninjam

import "fmt"

// AuthorizationFailed is an error thrown when the user cannot be
// connected to the NINJAM server
type AuthorizationFailed struct {
	User string
}

// Error displays the error for AuthorizationFailed
func (e *AuthorizationFailed) Error() string {
	return fmt.Sprintf("Authoriazation failed for user '%s'", e.User)
}

// CommandFailed is an error thrown when NINJAM responds with a non-zero
// response code.
type CommandFailed struct {
	Code int
}

// Error displays the error for CommandFailed
func (e *CommandFailed) Error() string {
	return fmt.Sprintf("Command Failed (NINJAM error code: %d)", e.Code)
}
