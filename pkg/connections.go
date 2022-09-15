package mpopenvidu

type Connections struct {
	Count   int64        `json:"numberOfElements"`
	Content []Connection `json:"content"`
}

type Connection struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}
