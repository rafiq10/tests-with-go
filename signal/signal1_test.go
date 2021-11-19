package signal

import (
	"testing"
)

func TestHandler(t *testing.T) {
	t.Log("plain log message")             // works like fmt.Print
	t.Logf("message with number: %d", 123) // works like fmt.Printf

}
