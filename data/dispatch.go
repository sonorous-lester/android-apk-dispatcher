package data

type DispatchBody struct {
	EventType string `json:"event_type"`
	ClientPayload ClientPayload `json:"client_payload"`
}

type ClientPayload struct {
	Variant   string `json:"variant"`
	Branch    string `json:"branch"`
	ChannelId string `json:"channel_id"`
	UserId    string `json:"user_id"`
}
