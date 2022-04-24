package fmpcloud

import (
	"testing"
	"time"
)

func TestWebsocketClient(t *testing.T) {
	websocketClient, err := NewWebsocketClient(WebsocketConfig{
		APIKey:       "test",
		WebsocketUrl: WebsocketCrypto,
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

	if err := websocketClient.Subscribe("btcusd"); err != nil {
		t.Fatal(err)
	}

	time.Sleep(30 * time.Second)
	if err := websocketClient.Close(); err != nil {
		t.Fatal(err)
	}

	t.Log("Success")
}
