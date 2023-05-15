package structures

type Chat struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	Target  *Room  `json:"target"`
}

func NewChat() *Chat {
	return &Chat{}
}

func (m *Chat) Chat() {
	// Model1 için metodlar burada yer alır
}
