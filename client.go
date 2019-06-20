package main

import (
	stan "github.com/nats-io/stan.go"
)

type NatsStreamingClient struct {
	Host      string
	Port      int
	ID        string
	ClusterID string
	conn      stan.Conn
}

func generateClientID() string {
	//TODO: 采用一个随机字符串防止冲突
	idSuffix := "foo"
	return "nats-streaming_cli_" + idSuffix
}

func (n *NatsStreamingClient) Publish(channelName string, content []byte) error {
	if n.conn == nil {
		opts := []stan.Option{}
		conn, err := stan.Connect(n.ClusterID, n.ID, opts...)
		if err != nil {
			return err
		}
		n.conn = conn
	}
	client := n.conn
	err := client.Publish(channelName, content)
	return err
}
