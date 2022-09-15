package mpopenvidu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Api struct {
	endpointUrl string
	secret      string
	Client      HttpClient
}

func (api *Api) GetSessions() Sessions {
	var sessions Sessions
	req, err := http.NewRequest("GET", api.endpointUrl+"sessions", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("OPENVIDUAPP", api.secret)
	res, err := api.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &sessions)
	return sessions
}

func (api *Api) GetConnections(sessions Sessions) int64 {
	var connections int64 = 0
	for _, session := range sessions.Content {
		connections += session.Connections.Count
	}
	return connections
}

func (api *Api) GetMediaNodes() MediaNodes {
	var mediaNodes MediaNodes
	req, err := http.NewRequest("GET", api.endpointUrl+"media-nodes", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("OPENVIDUAPP", api.secret)
	res, err := api.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &mediaNodes)
	if err != nil {
		log.Fatal(err)
	}
	return mediaNodes
}
