package main

import (
	"fmt"
	"github.com/xieqiaoyu/nats-streaming-cli/cmd"
)

func ShowCmd() *cmd.SubCmdResolver {
	return cmd.NewSubCmdResolver(cmd.CmdMap{
		"channel":  ShowChannelCmd,
		"channels": ShowChannelsCmd,
		"server":   ShowServerCmd,
		"store":    ShowStoreCmd,
		"clients":  ShowClientsCmd,
	})
}

func ShowChannelCmd(t ...string) {
	if len(t) < 1 {
		fmt.Println("Usage: show channel CHANNEL")
		return
	}
	info, err := monitor.GetChannelInfo(t[0], false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(info))
	}
}

func ShowChannelsCmd(t ...string) {
	info, err := monitor.GetChannelsInfo(false, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(info))
	}
}

func ShowServerCmd(t ...string) {
	info, err := monitor.GetServerInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(info))
	}
}

func ShowStoreCmd(t ...string) {
	info, err := monitor.GetStoreInfo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(info))
	}

}

func ShowClientsCmd(t ...string) {
	info, err := monitor.GetClientsInfo(false, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(info))
	}
}
