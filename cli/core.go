package cli

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	nclient "github.com/xieqiaoyu/nats-streaming-cli/client"
	"strings"
)

var (
	client  *nclient.NatsStreamingClient
	monitor *nclient.NatsStreamingMonitor
	rootCmd = RootCmd()
)

func livePrefix() (string, bool) {
	return "", false
}

func executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	rootCmd.Resolve(blocks...)
}

func completer(in prompt.Document) []prompt.Suggest {
	input := in.TextBeforeCursor()
	if input == "" {
		return []prompt.Suggest{}
	}
	blocks := strings.Split(input, " ")
	return rootCmd.Complete(blocks...)
}
func Run() {
	var port = 4222
	var httpPort = 8222
	var host = "localhost"
	var clientID = ""
	var clusterID = ""
	var err error

	monitor = &nclient.NatsStreamingMonitor{
		Host:     host,
		HttpPort: httpPort,
	}
	client = &nclient.NatsStreamingClient{
		Host: host,
		Port: port,
	}
	if clientID == "" {
		clientID = nclient.GenerateClientID()
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
