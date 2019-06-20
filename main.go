package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var server *NatsStreaming

func livePrefix() (string, bool) {
	return "", false
}

func badCmd() {
	fmt.Println("Sorry, I don't understand.")
}

func executor(in string) {
	var err error
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "show":
		var info string
		switch blocks[1] {
		case "channels":
			info, err = server.GetChannelsInfo()
		case "server":
			info, err = server.GetServerInfo()
		case "store":
			info, err = server.GetStoreInfo()
		case "clients":
			info, err = server.GetClientsInfo()
		default:
			badCmd()
			return
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(info)
		}
		return
	case "exit":
		os.Exit(0)
	default:
	}
	badCmd()
	return
}

func main() {
	var port = 4222
	var httpPort = 8222
	var host = "localhost"
	server = &NatsStreaming{
		Host:     host,
		Port:     port,
		HttpPort: httpPort,
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("[nats-streaming] "+server.Host+" > "),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("nats-streaming-cli"),
	)
	p.Run()
}
