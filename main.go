package main

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var client *NatsStreamingClient
var monitor *NatsStreamingMonitor

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
		var info []byte
		switch blocks[1] {
		case "channel":
			if len(blocks) < 3 {
				fmt.Println("Usage: show channel CHANNEL")
				return
			}
			info, err = monitor.GetChannelInfo(blocks[2])
		case "channels":
			info, err = monitor.GetChannelsInfo()
		case "server":
			info, err = monitor.GetServerInfo()
		case "store":
			info, err = monitor.GetStoreInfo()
		case "clients":
			info, err = monitor.GetClientsInfo()
		default:
			badCmd()
			return
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(info))
		}
		return
	case "pub":
		if len(blocks) != 3 {
			fmt.Println("Usage: pub CHANNEL MSG")
			return
		}
		channelName := blocks[1]
		//TODO support message with space
		msg := blocks[2]
		err := client.Publish(channelName, []byte(msg))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Success!")
		}
		return
	case "list":
		if len(blocks) != 2 {
			fmt.Println("Usage: list CHANNEL")
			return
		}
		channelName := blocks[1]
		data, err := client.List(channelName, 0, 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", strings.Join(data, "\n"))
		}
		return
	case "test":
		//msg, err := client.GetEndOfDayMsg("test.temp")
		//if err != nil {
		//	fmt.Println(err)
		//} else {
		//	fmt.Printf("%s,%d", string(msg.Data), msg.Sequence)
		//}
		//return
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
	var clientID = ""
	var clusterID = ""
	var err error

	monitor = &NatsStreamingMonitor{
		Host:     host,
		HttpPort: httpPort,
	}
	client = &NatsStreamingClient{
		Host: host,
		Port: port,
	}
	if clientID == "" {
		clientID = generateClientID()
	}
	client.ID = clientID
	if clusterID == "" {
		clusterID, err = monitor.GetClusterID()
		if err != nil || clusterID == "" {
			fmt.Printf("Fail to Get clusterID,please make sure server enable monitor or set cluster id manual:%s\n", err)
			return
		}
	}
	client.ClusterID = clusterID

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("[nats-streaming] "+host+" > "),
		prompt.OptionLivePrefix(livePrefix),
		prompt.OptionTitle("nats-streaming-cli"),
	)
	p.Run()
}
