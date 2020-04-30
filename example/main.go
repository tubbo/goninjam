package main

import (
	"fmt"
	"os"

	"github.com/tubbo/goninjam"
)

func main() {
	ninjam := goninjam.ConnectAnonymously("jam.mindbrainmusic.com:2049", "foo")
	err := ninjam.Chat("hello world")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
