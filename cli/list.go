package cli

import (
	"fmt"
	"strconv"
	"strings"
)

//ListCmd  handel cmd "list"
func ListCmd(t ...string) {
	var err error
	var start, limit uint64
	var channelName string
	switch len(t) {
	case 3:
		limit, err = strconv.ParseUint(t[2], 10, 64)
		if err != nil {
			fmt.Println("limit must be a nonnegative integer")
			return
		}
		fallthrough
	case 2:
		start, err = strconv.ParseUint(t[1], 10, 64)
		if err != nil {
			fmt.Println("start must be a nonnegative integenr")
			return
		}
		fallthrough
	case 1:
		channelName = t[0]
	default:
		fmt.Println("Usage: list CHANNEL [START [LIMIT]]")
		return
	}
	data, err := client.List(channelName, start, limit)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", strings.Join(data, "\n"))
	}
	return
}
