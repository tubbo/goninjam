# go-ninjam

A [NINJAM][] client for Go.

## Usage

```go
package main

import "github.com/tubbo/goninjam"

func main() {
  ninjam := goninjam.Connect("localhost:2049", "username", "password")

  // Send a chat message
  ninjam.Chat("hello world")

  // Change the topic
  ninjam.Topic("new topic")

  // Change the BPM
  ninjam.BPM(128)

  // Change the BPI
  ninjam.BPI(16)

  // Kick a user
  ninjam.Kick("troll")
}
```
