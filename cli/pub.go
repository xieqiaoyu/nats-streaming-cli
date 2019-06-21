package cli

import (
	"fmt"
)

func PubCmd(t ...string) {
	if len(t) != 2 {
		fmt.Println("Usage: pub CHANNEL MSG")
		return
	}
	channelName := t[0]
	//TODO support message with space
	msg := t[1]
	err := client.Publish(channelName, []byte(msg))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
	return
}
