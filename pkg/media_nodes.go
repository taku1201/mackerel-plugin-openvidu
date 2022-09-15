package mpopenvidu

type MediaNodes struct {
	Count   int64       `json:"numberOfElements"`
	Content []MediaNode `json:"content"`
}

type MediaNode struct {
	ID             string  `json:"id"`
	Object         string  `json:"object"`
	IP             string  `json:"ip"`
	Uri            string  `json:"uri"`
	Connected      bool    `json:"connected"`
	ConnectionTime int64   `json:"connectionTime"`
	Load           float32 `json:"load"`
	EnvironmentID  string  `json:"environmentId"`
	Status         string  `json:"status"`
	LaunchingTime  int64   `json:"launchingTime"`
}
