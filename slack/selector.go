package slack

import (
	"apkDispatcher/data"
	"apkDispatcher/env"
)

type Selector struct {
	Channel     string       `json:"channel"`
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Text           string   `json:"text"`
	CallbackId     string   `json:"callback_id"`
	Color          string   `json:"color"`
	AttachmentType string   `json:"attachment_type"`
	Actions        []data.Action `json:"actions"`
}

func CreateBuildVariantBtnSelector(channel string) Selector {
	actions := mapVariantsToAction(env.GetVariant())
	return Selector{
		Channel: channel,
		Text:    "Please select a variant",
		Attachments: []Attachment{
			{
				Text:           "Select a variant",
				CallbackId:     "select_variant",
				Color:          "#3AA3E3",
				AttachmentType: "default",
				Actions: actions,
			},
		},
	}
}

func mapVariantsToAction(variants []string) []data.Action {
	actions := make([]data.Action, len(variants))
	for i, variant := range variants {
		actions[i] = data.Action{
			Name:            "variant",
			Text:            variant,
			Type:            "button",
			Value:           variant,
			SelectedOptions: nil,
			Options:         nil,
		}
	}
	return actions
}

func CreateBranchMenuMsgSelector(channel string, options []data.Option) Selector {
	return Selector{
		Channel: channel,
		Text:    "Please select a branch",
		Attachments: []Attachment{
			{
				Text:           "Select a branch",
				CallbackId:     "select_branch",
				Color:          "#3AA3E3",
				AttachmentType: "default",
				Actions: []data.Action{
					{
						Name:    "branch",
						Text:    "--Branch--",
						Type:    "select",
						Options: options,
					},
				},
			},
		},
	}
}
