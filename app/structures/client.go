package structures

import "github.com/gofiber/websocket/v2"

type Client struct {
	Conn     *websocket.Conn
	Username string `json:"username"`
	RoomID   string `json:"-"`
}

func NewClient() *Client {
	return &Client{}
}

func (m *Client) Client() {
	// Model1 için metodlar burada yer alır
}

// Chat Mesajı
type StatusMessage struct {
	Message string `json:"message"`
}

func (p *StatusMessage) TableName() string {
	return "calendar.users"
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Jwt     string      `json:"jwt"`
	Data    interface{} `json:"data"`
}

type AuthErrorResponse struct {
	Message string `json:"message"`
}
