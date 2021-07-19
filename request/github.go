package request

import (
	"apkDispatcher/data"
	"apkDispatcher/storage"
	"log"
	"time"
)

type branch struct {
	Name string `json:"name"`
}

func (req *request) FetchBranches() []data.Option {
	res, err := req.client.R().
		SetHeaders(map[string]string{
			"Accept":        "application/vnd.github.v3+json",
			"Authorization": "token " + req.token,
		}).
		SetResult([]branch{}).
		Get(req.branchUrl)

	if err != nil {
		log.Printf("fetchBranch Error %v \n", err)
		return []data.Option{}
	}

	if res.IsError() {
		log.Printf("fetch repo branch error: %v \n", res.Error())
		return []data.Option{}
	}

	result := res.Result().(*[]branch)
	var options []data.Option
	for _, branch := range *result {
		options = append(options, data.Option{
			Text:  branch.Name,
			Value: branch.Name,
		})
	}
	return options
}
func (req *request) TriggerDispatch(payload *data.ActionPayload) bool{
	channelId := payload.Channel.Id
	userId := payload.User.Id
	variant := payload.Actions[0].Value
	store := storage.GetInstance()
	branch := store.GetBranch(userId)
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	_, err := req.client.R().
		SetHeader("Authorization", "Bearer "+req.token).
		SetBody(&data.DispatchBody{
			EventType:     currentTime,
			ClientPayload: data.ClientPayload{
				Variant:   variant,
				Branch:    branch,
				ChannelId: channelId,
				UserId:    userId,
			},
		}).Post(req.dispatchUrl)
	if err != nil {
		log.Printf("send dispatch goes wrong %v", err)
		return false
	}
	return true
}