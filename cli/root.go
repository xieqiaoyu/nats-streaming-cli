package cli

import (
	"fmt"
	"github.com/xieqiaoyu/nats-streaming-cli/cmd"
	"os"
)

func RootCmd() *cmd.SubCmdResolver {
	cmdMap := cmd.CmdMap{
		"show": ShowCmd(),
		"pub":  PubCmd,
		"list": ListCmd,
		"help": HelpCmd,
		"?":    HelpCmd,
		"exit": ExitCmd,
		"test": TestCmd,
	}
	suggestions := cmd.Suggest{
		"show": "show server mesage",
		"pub":  "publish msg to a channel",
		"list": "list message in a channel",
		"help": "show help message",
		"exit": "Exit cli",
	}
	return cmd.NewSubCmdResolver(cmdMap, suggestions)
}

func ExitCmd(t ...string) {
	os.Exit(0)
}

func TestCmd() {
	fmt.Println("test is a command for development test")
}

const helpStr = `
show
  ├─ channel CHANNEL                show a specific channel info
  ├─ channels                       show all channel info
  ├─ server                         show server info
  ├─ store                          show store info
  └─ clients                        show clients info

pub CHANNEL MESSAGE                 publish MESSAGE to CHANNEL

list CHANNEL [START [LIMIT]]        list LIMIT(unlimit if not specific) messages in CHANNEL start at START (0 if not specific)

help                                print this help message

exit                                exit
`

func HelpCmd() {
	fmt.Println(helpStr)
}
