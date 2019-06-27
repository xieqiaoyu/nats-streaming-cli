package cli

import (
	"fmt"
	"strings"
)

func PubCmd(t ...string) {
	if len(t) < 2 {
		fmt.Println("Usage: pub CHANNEL MESSAGE")
		return
	}
	channelName := t[0]
	msg := strings.Join(t[1:], " ")
	err := client.Publish(channelName, []byte(msg))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
	return
}
