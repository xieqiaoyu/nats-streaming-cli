package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type NatsStreaming struct {
	Host      string
	Port      int
	HttpPort  int
	HttpsPort int
	UseHttps  bool
}

func (n *NatsStreaming) fetchMonitorEndpoint(endpoint string, params ...[]string) (string, error) {
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
		query := new(url.Values)
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
		return "", err
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (n *NatsStreaming) GetChannelsInfo() (string, error) {
	return n.fetchMonitorEndpoint("channelsz")
}

func (n *NatsStreaming) GetServerInfo() (string, error) {
	return n.fetchMonitorEndpoint("serverz")
}

func (n *NatsStreaming) GetStoreInfo() (string, error) {
	return n.fetchMonitorEndpoint("storez")
}

func (n *NatsStreaming) GetClientsInfo() (string, error) {
	return n.fetchMonitorEndpoint("clientsz")
}
