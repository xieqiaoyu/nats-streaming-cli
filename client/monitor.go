package client

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

//NatsStreamingMonitor client to access nats-streaming http monitor service
type NatsStreamingMonitor struct {
	Host      string
	HTTPPort  int
	HTTPSPort int
	UseHTTPS  bool
}

func (n *NatsStreamingMonitor) fetchMonitorEndpoint(endpoint string, params ...[]string) ([]byte, error) {
	callURL := url.URL{
		Path: path.Join("streaming", endpoint),
	}
	if n.UseHTTPS {
		callURL.Scheme = "https"
		callURL.Host = net.JoinHostPort(n.Host, strconv.Itoa(n.HTTPSPort))
	} else {
		callURL.Scheme = "http"
		callURL.Host = net.JoinHostPort(n.Host, strconv.Itoa(n.HTTPPort))
	}
	if len(params) > 0 {
		query := url.Values{}
		for _, value := range params {
			switch len(value) {
			case 0:
			case 1:
				query.Set(value[0], "")
			default:
				query.Set(value[0], value[1])
			}
		}
		callURL.RawQuery = query.Encode()
	}
	req, err := http.NewRequest("GET", callURL.String(), nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//GetChannelInfo get the channel info
func (n *NatsStreamingMonitor) GetChannelInfo(channelName string, subs bool) ([]byte, error) {
	params := [][]string{
		{"channel", channelName},
	}
	if subs {
		params = append(params, []string{"subs", "1"})
	}
	return n.fetchMonitorEndpoint("channelsz", params...)
}

//GetChannelsInfo get batch channels info
func (n *NatsStreamingMonitor) GetChannelsInfo(subs bool, offset, limit uint64) ([]byte, error) {
	params := [][]string{}
	if subs {
		params = append(params, []string{"subs", "1"})
	}
	if offset > 0 {
		params = append(params, []string{"offset", strconv.FormatUint(offset, 10)})
	}
	if limit > 0 {
		params = append(params, []string{"limit", strconv.FormatUint(limit, 10)})
	}
	return n.fetchMonitorEndpoint("channelsz", params...)
}

//GetServerInfo get server info
func (n *NatsStreamingMonitor) GetServerInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("serverz")
}

//GetStoreInfo get store info
func (n *NatsStreamingMonitor) GetStoreInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("storez")
}

//GetClientInfo get info of a client
func (n *NatsStreamingMonitor) GetClientInfo(clientName string, subs bool) ([]byte, error) {
	params := [][]string{
		{"client", clientName},
	}
	if subs {
		params = append(params, []string{"subs", "1"})
	}
	return n.fetchMonitorEndpoint("clientsz", params...)
}

//GetClientsInfo get batch client info
func (n *NatsStreamingMonitor) GetClientsInfo(subs bool, offset, limit uint64) ([]byte, error) {
	params := [][]string{}
	if subs {
		params = append(params, []string{"subs", "1"})
	}
	if offset > 0 {
		params = append(params, []string{"offset", strconv.FormatUint(offset, 10)})
	}
	if limit > 0 {
		params = append(params, []string{"limit", strconv.FormatUint(limit, 10)})
	}
	return n.fetchMonitorEndpoint("clientsz", params...)
}

//Serverz struct of serverz endpoint return data
type Serverz struct {
	ClusterID     string `json:"cluster_id"`
	State         string `json:"state"`
	Uptime        string `json:"uptime"`
	Role          string `json:"role"`
	Clients       uint32 `json:"clients"`
	Channels      uint32 `json:"channels"`
	Subscriptions uint32 `json:"subscriptions"`
}

//GetClusterID fetch cluster from nats-streaming server
func (n *NatsStreamingMonitor) GetClusterID() (string, error) {
	serverInfo, err := n.GetServerInfo()
	if err != nil {
		return "", err
	}
	serverz := new(Serverz)
	err = json.Unmarshal(serverInfo, serverz)
	if err != nil {
		return "", err
	}
	return serverz.ClusterID, nil
}
