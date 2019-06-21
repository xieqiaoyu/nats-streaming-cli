package main

import (
	"fmt"
	"github.com/xieqiaoyu/nats-streaming-cli/cmd"
	"os"
)

func RootCmd() *cmd.SubCmdResolver {
	return cmd.NewSubCmdResolver(cmd.CmdMap{
		"show": ShowCmd(),
		"pub":  PubCmd,
		"list": ListCmd,
		"test": TestCmd,
		"exit": ExitCmd,
	})
}

func ExitCmd(t ...string) {
	os.Exit(0)
}

func TestCmd() {
	fmt.Println("Run Test!")
}
