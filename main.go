package main

import (
	"apkDispatcher/data"
	"apkDispatcher/env"
	"apkDispatcher/request"
	"apkDispatcher/slack"
	"apkDispatcher/storage"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	ownerName := os.Getenv("OWNER_NAME")
	projectName := os.Getenv("PROJECT_NAME")
	variants := generateVariants(os.Getenv("VARIANT"))
	token := os.Getenv("TOKEN")
	env.UpdateEnv(ownerName, projectName, token, variants)

	slashEndPoint := os.Getenv("SLASH_ENDPOINT")
	interactivityEndPoint := os.Getenv("INTERACTIVITY_ENDPOINT")
	port := os.Getenv("PORT")

	http.HandleFunc(slashEndPoint, askBranch)
	http.HandleFunc(interactivityEndPoint, action)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatalf("listen server goes wrong: %v \n", err)
	}
}

func generateVariants(variant string) []string {
	var variants []string
	if len(variant) != 0 {
		variants = strings.Split(variant, ",")
	}else{
		variants = []string{"debug", "release"}
	}
	return variants
}

func askBranch(w http.ResponseWriter, r *http.Request) {
	query, _ := convertBodyToUrlQuery(r.Body)
	resUrl := query.Get("response_url")
	channelId := query.Get("channel_id")

	req := request.GetClient()
	options := req.FetchBranches()
	req.SendBranchMenuMessage(channelId, resUrl, options)

}

func convertBodyToUrlQuery(r io.Reader) (url.Values, error){
	body, _ := ioutil.ReadAll(r)
	strBody := string(body)
	decodeBody, _ := url.QueryUnescape(strBody)
	return url.ParseQuery(decodeBody)
}

func action(w http.ResponseWriter, r *http.Request) {
	actionPayload := data.ConvertBodyToActionPayload(r.Body)
	callbackId := actionPayload.CallbackId
	req := request.GetClient()

	if callbackId == "select_branch" {
		storeSelectedBranchName(actionPayload)
		selector := slack.CreateBuildVariantBtnSelector(actionPayload.CallbackId)
		req.SendVariantBtnMessage(selector, actionPayload.ResponseUrl)
	}

	if callbackId == "select_variant" {
		responseUrl := actionPayload.ResponseUrl
		dispatchSuccess := req.TriggerDispatch(actionPayload)
		if dispatchSuccess {
			req.SendMessage(responseUrl)
			removeSelectedBranch(actionPayload)
		}
	}
}

var store = storage.GetInstance()
func storeSelectedBranchName(payload *data.ActionPayload){
	userId := payload.User.Id
	versionValue := payload.Actions[0].SelectedOptions[0].Value
	store.Add(userId, versionValue)
}

func removeSelectedBranch(payload *data.ActionPayload){
	store.Remove(payload.User.Id)
}

