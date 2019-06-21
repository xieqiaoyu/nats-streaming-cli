package client

import (
	"fmt"
	stan "github.com/nats-io/stan.go"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
)

type NatsStreamingClient struct {
	Host      string
	Port      int
	ID        string
	ClusterID string
	conn      stan.Conn
}

func GenerateClientID() string {
	// generate random string
	// reference  https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	src := rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(10)
	dice := src.Int63()
	for i := 0; i <= 9; i++ {
		idx := int(dice & letterIdxMask)
		sb.WriteByte(letterBytes[idx])
		dice >>= letterIdxBits
	}

	idSuffix := sb.String()
	return "nats-streaming_cli_" + idSuffix
}

func (n *NatsStreamingClient) getConn() (stan.Conn, error) {
	if n.conn == nil {
		opts := []stan.Option{}
		conn, err := stan.Connect(n.ClusterID, n.ID, opts...)
		if err != nil {
			return nil, err
		}
		n.conn = conn
	}
	return n.conn, nil

}

func (n *NatsStreamingClient) Publish(channelName string, content []byte) error {
	client, err := n.getConn()
	if err != nil {
		return err
	}
	return client.Publish(channelName, content)
}

func (n *NatsStreamingClient) Close() {
	if n.conn != nil {
		n.Close()
	}
}

func (n *NatsStreamingClient) List(channelName string, startAt, limit uint64) ([]string, error) {
	endOfDay, err := n.GetEndOfDayMsg(channelName)
	if err != nil {
		return nil, err
	}
	sequence := endOfDay.Sequence

	client, err := n.getConn()
	if err != nil {
		return nil, err
	}
	subOpts := []stan.SubscriptionOption{
		stan.DeliverAllAvailable(),
	}

	if startAt > 0 {
		subOpts = append(subOpts, stan.StartAtSequence(startAt))
	}

	msgChan := make(chan *stan.Msg, 10)
	defer close(msgChan)
	var stop bool

	sub, err := client.Subscribe(channelName, func(msg *stan.Msg) {
		if !stop {
			msgChan <- msg
		}
	}, subOpts...)

	if err != nil {
		return nil, err
	}
	defer sub.Close()

	result := []string{}
	var count uint64
	for {
		m := <-msgChan
		result = append(result, string(m.Data))
		if m.Sequence == sequence {
			stop = true
			break
		}
		if limit > 0 {
			count++
			if count >= limit {
				stop = true
				break
			}
		}
	}
	return result, nil
}

func (n *NatsStreamingClient) GetEndOfDayMsg(channelName string) (*stan.Msg, error) {
	client, err := n.getConn()
	if err != nil {
		return nil, err
	}
	subOpts := []stan.SubscriptionOption{
		stan.StartWithLastReceived(),
	}
	msgChan := make(chan *stan.Msg, 1)
	defer close(msgChan)
	var gocha bool

	sub, err := client.Subscribe(channelName, func(msg *stan.Msg) {
		if !gocha {
			msgChan <- msg
			gocha = true
		}
	}, subOpts...)

	if err != nil {
		return nil, err
	}
	defer sub.Close()

	var msg *stan.Msg
	select {
	case m := <-msgChan:
		msg = m
	case <-time.After(2 * time.Second):
		return nil, fmt.Errorf("Get msg time out,maybe the channel is Empty")
	}
	return msg, nil
}
