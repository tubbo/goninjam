package goninjam

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"
)

const MESSAGE_SERVER_AUTH_CHALLENGE = 0x00
const MESSAGE_SERVER_AUTH_REPLY = 0x01
const MESSAGE_SERVER_CONFIG_CHANGE_NOTIFY = 0x02
const MESSAGE_SERVER_USERINFO_CHANGE_NOTIFY = 0x03
const MESSAGE_SERVER_DOWNLOAD_INTERVAL_BEGIN = 0x04
const MESSAGE_SERVER_DOWNLOAD_INTERVAL_WRITE = 0x05
const MESSAGE_CLIENT_AUTH_USER = 0x80
const MESSAGE_CLIENT_SET_USERMASK = 0x81
const MESSAGE_CLIENT_SET_CHANNEL_INFO = 0x82
const MESSAGE_CLIENT_UPLOAD_INTERVAL_BEGIN = 0x83
const MESSAGE_CLIENT_UPLOAD_INTERVAL_WRITE = 0x84
const MESSAGE_CHAT_MESSAGE = 0xC0

// Client is the NINJAM client instance connected to the server
type Client struct {
	Host string
	User string
}

func prepend(x []byte, y byte) []byte {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

// Send socket data over the TCP socket to NINJAM
func (c *Client) Send(method byte, arguments []byte) error {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", c.Host)

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	payload := prepend(arguments, method)
	response, err := conn.Write(payload)

	if err != nil {
		return err
	}

	if response != 0 {
		return &CommandFailed{Code: response}
	}

	return nil
}

// Chat sends a message to the chat
func (c *Client) Chat(message string) error {
	return c.Send(MESSAGE_CHAT_MESSAGE, []byte(message))
}

func combine(slice []byte, addition string) []byte {
	for _, b := range []byte(addition) {
		slice = append(slice, b)
	}

	return slice
}

// Topic updates the topic in the server
func (c *Client) Topic(text string) error {
	var args []byte

	args = combine(args, "TOPIC ")
	args = combine(args, text)

	return c.Send(MESSAGE_CHAT_MESSAGE, args)
}

// Kick a user from the server
func (c *Client) Kick(user string) error {
	var args []byte

	args = combine(args, user)
	args = combine(args, " K")

	return c.Send(MESSAGE_CLIENT_SET_USERMASK, args)
}

// BPM sets the current tempo of the channel on the next loop
func (c *Client) BPM(bpm int) error {
	var args []byte

	args = combine(args, "BPM ")
	args = combine(args, strconv.Itoa(bpm))

	return c.Send(MESSAGE_CLIENT_SET_CHANNEL_INFO, args)
}

// BPI sets the amount of beats in each interval on the next loop
func (c *Client) BPI(bpi int) error {
	var args []byte

	args = combine(args, "BPI ")
	args = combine(args, strconv.Itoa(bpi))

	return c.Send(MESSAGE_CLIENT_SET_CHANNEL_INFO, args)
}

// Authorize uses a given password to connect to the NINJAM server.
func (c *Client) Authorize(password string) error {
	var args []byte

	args = combine(args, c.User)
	args = combine(args, " ")
	args = combine(args, password)

	return c.Send(MESSAGE_CLIENT_AUTH_USER, args)
}
