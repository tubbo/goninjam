package goninjam

import "fmt"

// Connect establishes a connection with the NINJAM server over TCP. It
// requires a username and password to log in. You can also use
// ConnectAnonymously to refrain from providing a password.
func Connect(host string, username string, password string) (Client, error) {
	var err error
	var client Client

	if password == "" {
		return ConnectAnonymously(host, username), nil
	}

	client = Client{Host: host, User: username}
	err = client.Authorize(password)

	return client, err
}

// ConnectAnonymously works just like Connect, but does not include a
// password and authorization step. It also prepends "anonymous:" to the
// username as that's what the NINJAM server is expecting for anonymous
// users.
func ConnectAnonymously(host string, username string) Client {
	anonymous := fmt.Sprintf("anonymous:%s", username)

	return Client{Host: host, User: anonymous}
}
