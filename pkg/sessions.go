package mpopenvidu

type Sessions struct {
	Count   int64     `json:"numberOfElements"`
	Content []Session `json:"content"`
}

type Session struct {
	ID                     string      `json:"id"`
	Object                 string      `json:"object"`
	SessionID              string      `json:"sessionId"`
	CreatedAt              int64       `json:"createdAt"`
	MediaMode              string      `json:"mediaMode"`
	RecordingMode          string      `json:"recordingMode"`
	DefaultOutputMode      string      `json:"defaultOutputMode"`
	DefaultRecordingLayout string      `json:"defaultRecordingLayout"`
	CustomSessionID        string      `json:"customSessionId"`
	Connections            Connections `json:"connections"`
	Recording              bool        `json:"recording"`
	ForcedVideoCodec       string      `json:"forcedVideoCodec"`
	AllowTranscoding       bool        `json:"allowTranscoding"`
	MediaNodeID            string      `json:"mediaNodeId"`
}
