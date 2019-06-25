package cli

import (
	"fmt"
	"strings"
)

func PubCmd(t ...string) {
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
