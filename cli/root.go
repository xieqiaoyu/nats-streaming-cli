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
		"test": TestCmd,
		"exit": ExitCmd,
	}
	suggestions := cmd.Suggest{
		"show": "show server mesage",
		"pub":  "publish msg to a channel",
		"list": "list message in a channel",
		"exit": "Exit cli",
	}
	return cmd.NewSubCmdResolver(cmdMap, suggestions)
}

func ExitCmd(t ...string) {
	os.Exit(0)
}

func TestCmd() {
	fmt.Println("Run Test!")
}
