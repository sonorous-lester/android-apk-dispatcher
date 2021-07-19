package request

import (
	"apkDispatcher/env"
	"github.com/go-resty/resty/v2"
)

const githubDomain = "https://api.github.com/repos/"
const (
	Branch = "/branches"
	Dispatcher = "/dispatches"
)

type request struct {
	branchUrl string
	dispatchUrl string
	token string
	client *resty.Client
}

func GetClient() *request{
	owner := env.GetOwnerName()
	project := env.GetProjectName()
	token := env.GetToken()
	return &request{
		branchUrl: generateUrl(owner, project, Branch),
		dispatchUrl: generateUrl(owner, project, Dispatcher),
		token: token,
		client: resty.New(),
	}
}

func generateUrl(ownerName string, project string, endPoint string) string{
	return githubDomain + ownerName + "/" + project + endPoint
}

