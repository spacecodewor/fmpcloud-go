package fmpcloud

import (
	"testing"
	"time"
)

func TestWebsocketClient(t *testing.T) {
	websocketClient, err := NewWebsocketClient(WebsocketConfig{
		APIKey:       "demo",
		WebsocketUrl: WebsocketStock,
		PingPongCfg: &PingPongConfig{
			IsEnabled:  true,
			PongWait:   DefaultPongWait,
			PingPeriod: DefaultPingPeriod,
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	go func() {
		err = websocketClient.RunReadLoop(func(event interface{}) error {
			t.Log("Got event: ", event)
			return nil
		})
		if err != nil {
			t.Log("FATAL: ", err)
			return
		}
	}()

	if err := websocketClient.Login(); err != nil {
		t.Fatal(err)
	}

	// if err := websocketClient.Subscribe("btcusd"); err != nil {
	// 	t.Fatal(err)
	// }

	if err := websocketClient.Subscribe("aapl"); err != nil {
		t.Fatal(err)
	}

	if err := websocketClient.Subscribe("amd"); err != nil {
		t.Fatal(err)
	}

	if err := websocketClient.Subscribe("fb"); err != nil {
		t.Fatal(err)
	}

	if err := websocketClient.Subscribe("twtr"); err != nil {
		t.Fatal(err)
	}

	time.Sleep(60 * time.Second)
	// if err := websocketClient.Unsubscribe("btcusd"); err != nil {
	// 	t.Fatal(err)
	// }
	if err := websocketClient.Unsubscribe("aapl"); err != nil {
		t.Fatal(err)
	}
	if err := websocketClient.Unsubscribe("amd"); err != nil {
		t.Fatal(err)
	}
	if err := websocketClient.Unsubscribe("fb"); err != nil {
		t.Fatal(err)
	}
	if err := websocketClient.Unsubscribe("twtr"); err != nil {
		t.Fatal(err)
	}

	time.Sleep(10 * time.Second)
	if err := websocketClient.Close(); err != nil {
		t.Fatal(err)
	}

	t.Log("Success")
}
