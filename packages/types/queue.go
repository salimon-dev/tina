package types

type BaseEvent struct {
	Action string `json:"action"`
}

type ThreadEvent struct {
	Action string `json:"action"`
	Type   string `json:"type"`
	Thread Thread `json:"thread"`
}

type MessageEvent struct {
	Action  string  `json:"action"`
	Message Message `json:"message"`
}

type TransactionEvent struct {
	Action      string      `json:"action"`
	Transaction Transaction `json:"transaction"`
	Type        string      `json:"type"`
}
