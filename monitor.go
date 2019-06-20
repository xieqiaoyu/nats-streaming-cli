package main

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type NatsStreamingMonitor struct {
	Host      string
	HttpPort  int
	HttpsPort int
	UseHttps  bool
}

func (n *NatsStreamingMonitor) fetchMonitorEndpoint(endpoint string, params ...[]string) ([]byte, error) {
	callURL := url.URL{
		Path: path.Join("streaming", endpoint),
	}
	if n.UseHttps {
		callURL.Scheme = "https"
		callURL.Host = net.JoinHostPort(n.Host, strconv.Itoa(n.HttpsPort))
	} else {
		callURL.Scheme = "http"
		callURL.Host = net.JoinHostPort(n.Host, strconv.Itoa(n.HttpPort))
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

func (n *NatsStreamingMonitor) GetChannelInfo(channelName string) ([]byte, error) {
	params := [][]string{
		[]string{"channel", channelName},
	}
	return n.fetchMonitorEndpoint("channelsz", params...)
}

func (n *NatsStreamingMonitor) GetChannelsInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("channelsz")
}

func (n *NatsStreamingMonitor) GetServerInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("serverz")
}

func (n *NatsStreamingMonitor) GetStoreInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("storez")
}

func (n *NatsStreamingMonitor) GetClientsInfo() ([]byte, error) {
	return n.fetchMonitorEndpoint("clientsz")
}

type Serverz struct {
	ClusterID     string `json:"cluster_id"`
	State         string `json:"state"`
	Uptime        string `json:"uptime"`
	Role          string `json:"role"`
	Clients       uint32 `json:"clients"`
	Channels      uint32 `json:"channels"`
	Subscriptions uint32 `json:"subscriptions"`
}

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
