package message

type Message struct {
	Type    string `json:"type"` // "public" or "private"
	Content string `json:"content"`
	Target  string `json:"target,omitempty"` // Target client ID for private messages
}
