package mpopenvidu

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

type OpenViduPlugin struct {
	prefix      string
	endpointUrl string
	secret      string
}

func GetSessions(endpointUrl string, secret string) Sessions {
	var sessions Sessions
	client := http.Client{}
	req, err := http.NewRequest("GET", endpointUrl+"sessions", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("OPENVIDUAPP", secret)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &sessions)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetConnections(sessions Sessions) int64 {
	var connections int64 = 0
	for _, session := range sessions.Content {
		connections += session.Connections.Count
	}
	return connections
}

func GetMediaNodes(endpointUrl string, secret string) MediaNodes {
	var mediaNodes MediaNodes
	client := http.Client{}
	req, err := http.NewRequest("GET", endpointUrl+"media-nodes", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("OPENVIDUAPP", secret)
	res, err := client.Do(req)
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

func (p OpenViduPlugin) FetchMetrics() (map[string]float64, error) {
	stat := make(map[string]float64)
	// sessions
	sessions := GetSessions(p.endpointUrl, p.secret)
	stat["sessions"] = float64(sessions.Count)
	// connections
	connections := GetConnections(sessions)
	stat["connections"] = float64(connections)
	// media nodes
	mediaNodes := GetMediaNodes(p.endpointUrl, p.secret)
	stat["media-nodes"] = float64(mediaNodes.Count)
	return stat, nil
}

func (p OpenViduPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(p.MetricKeyPrefix())
	var graphdef = map[string]mp.Graphs{
		"": {
			Label: labelPrefix,
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "sessions", Label: "sessions", Diff: false, Stacked: false},
				{Name: "connections", Label: "connections", Diff: false, Stacked: false},
				{Name: "media-nodes", Label: "media nodes", Diff: false, Stacked: false},
			},
		},
	}
	return graphdef
}

func (p OpenViduPlugin) MetricKeyPrefix() string {
	if p.prefix == "" {
		p.prefix = "openvidu"
	}
	return p.prefix
}

func Do() {
	var (
		optPrefix      = flag.String("metric-key-prefix", "", "Metric key prefix")
		optEndpointUrl = flag.String("endpoint-url", "", "API Endpoint URL")
		optSecret      = flag.String("secret", "", "OpenVidu Secret")
	)
	r := regexp.MustCompile("https?://(.*)/openvidu/api/?")
	flag.Parse()
	if *optSecret == "" {
		flag.Usage()
		fmt.Println("OpenVidu secret is empty")
		os.Exit(1)
	}
	if !r.MatchString(*optEndpointUrl) {
		flag.Usage()
		fmt.Println("Endpoint URL format is invalid")
		os.Exit(1)
	}
	if !strings.HasSuffix(*optEndpointUrl, "/") {
		*optEndpointUrl = *optEndpointUrl + "/"
	}
	mp.NewMackerelPlugin(&OpenViduPlugin{
		prefix:      *optPrefix,
		endpointUrl: *optEndpointUrl,
		secret:      *optSecret,
	}).Run()
}
