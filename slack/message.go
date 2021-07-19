package slack

type MessageBlock struct {
	Blocks []Message `json:"blocks"`
}

type Message struct {
	Type string      `json:"type"`
	Text TextSetting `json:"text"`
}

type TextSetting struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func CreateMessageBlock(message string) MessageBlock {
	return MessageBlock{
		Blocks: []Message{
			{
				Type: "section",
				Text: TextSetting{
					Type: "mrkdwn",
					Text: message,
				},
			},
		},
	}
}

