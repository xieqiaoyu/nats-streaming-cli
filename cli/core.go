package cli

import (
	"flag"
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	nclient "github.com/xieqiaoyu/nats-streaming-cli/client"
	"github.com/xieqiaoyu/nats-streaming-cli/metadata"
	"os"
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

const usageStr = `
Usage: nats-streaming-cli [options]
Options:
      -h   --host <string>           set nats-streaming server host (default: localhost)
      -p   --port <int>              set nats-streaming server port (default: 4222)
      -m   --http_port <int>         set http monitoring port (default: 8222)
      -cid --cluster_id <string>     set the server cluster ID, if not set, we will try to get cluster id from server monitor endpoint
           --client_id  <string>     specific client id cli use ,if not set, we will use a random client id
      -v   --version                 show version
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func Run() {
	fs := flag.NewFlagSet("nats-streaming-cli", flag.ExitOnError)
	fs.Usage = usage
	var port, httpPort uint
	var host, clientID, clusterID string
	fs.UintVar(&port, "p", 4222, "")
	fs.UintVar(&port, "port", 4222, "")

	fs.UintVar(&httpPort, "m", 8222, "")
	fs.UintVar(&httpPort, "http_port", 8222, "")

	fs.StringVar(&host, "h", "localhost", "")
	fs.StringVar(&host, "host", "localhost", "")

	fs.StringVar(&clientID, "client_id", "", "")

	fs.StringVar(&clusterID, "cid", "", "")
	fs.StringVar(&clusterID, "cluster_id", "", "")

	var showVersion bool
	fs.BoolVar(&showVersion, "v", false, "")
	fs.BoolVar(&showVersion, "version", false, "")

	fs.Parse(os.Args[1:])

	var err error
	if showVersion {
		fmt.Println(metadata.GetVersionString())
		return
	}

	monitor = &nclient.NatsStreamingMonitor{
		Host:     host,
		HttpPort: int(httpPort),
	}
	client = &nclient.NatsStreamingClient{
		Host: host,
		Port: int(port),
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
