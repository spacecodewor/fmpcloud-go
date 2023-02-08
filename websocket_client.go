package fmpcloud

import (
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Core params
const (
	WebsocketStock  WebsocketURL = "wss://websockets.financialmodelingprep.com"
	WebsocketCrypto WebsocketURL = "wss://crypto.financialmodelingprep.com"
	WebsocketForex  WebsocketURL = "wss://forex.financialmodelingprep.com"
)

// WebsocketURL type for websocket url
type WebsocketURL string

// WebsocketConfig for create new Websocket client
type WebsocketConfig struct {
	Logger *zap.Logger
	APIKey string
	URL    WebsocketURL
	Debug  bool
}

// WebsocketClient ...
type WebsocketClient struct {
	conn   *websocket.Conn
	apiKey string
	logger *zap.Logger
	debug  bool
}

// Event ...
type Event struct {
	Event   string   `json:"event"`
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Symbol  string   `json:"s"`
	Type    string   `json:"type"`
	Time    int64    `json:"t"`
	Ap      *float64 `json:"ap"`
	As      *float64 `json:"as"`
	Bp      *float64 `json:"bp"`
	Bs      *float64 `json:"bs"`
	Lp      *float64 `json:"lp"`
	Ls      *float64 `json:"ls"`
}

// NewWebsocketClient creates a new API client
func NewWebsocketClient(cfg WebsocketConfig) (*WebsocketClient, error) {
	websocketClient := &WebsocketClient{logger: cfg.Logger, debug: cfg.Debug, apiKey: cfg.APIKey}
	if websocketClient.logger == nil {
		logger, err := createNewLogger()
		if err != nil {
			return nil, errors.Wrap(err, "Error create new zap logger")
		}

		websocketClient.logger = logger
	}

	websocketClient.connect(cfg.URL)

	return websocketClient, nil
}

func (w *WebsocketClient) Close() error {
	return w.conn.Close()
}

func (w *WebsocketClient) connect(url WebsocketURL) {
	if w.conn != nil {
		return
	}

	ws, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return
	}

	w.conn = ws
}

func (w *WebsocketClient) Login() error {
	sub := `{"event": "login", "data": {"apiKey": "` + w.apiKey + `" }}`

	if err := w.conn.WriteMessage(websocket.TextMessage, []byte(sub)); err != nil {
		return errors.Wrap(err, "can't login in websocket server")
	}

	return nil
}

func (w *WebsocketClient) Subscribe(tiker string) error {
	sub := `{"event": "subscribe", "data": {"ticker": "` + tiker + `" }}`

	if err := w.conn.WriteMessage(websocket.TextMessage, []byte(sub)); err != nil {
		return errors.Wrap(err, "can't subscribe to event")
	}

	return nil
}

func (w *WebsocketClient) Unsubscribe(tiker string) error {
	sub := `{"event": "unsubscribe", "data": {"ticker": "` + tiker + `" }}`

	if err := w.conn.WriteMessage(websocket.TextMessage, []byte(sub)); err != nil {
		return errors.Wrap(err, "can't unsubscribe from event")
	}

	return nil
}

func (w *WebsocketClient) RunReadLoop(fn func(event Event) error) error {
	for {
		_, msg, err := w.conn.ReadMessage()
		if err != nil {
			return errors.Wrap(err, "can't read message")
		}

		var event Event
		if err := jsoniter.Unmarshal(msg, &event); err != nil {
			w.logger.Error(
				"Can't unmarshal event",
				zap.Error(err),
				zap.Any("message", string(msg)),
			)

			continue
		}

		if err := fn(event); err != nil {
			return err
		}
	}
}

// String ...
func (w WebsocketURL) String() string {
	return string(w)
}
