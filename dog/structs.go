package dog

import "time"

type LogResponse struct {
	Data []AttributeHolder `json:"data"`
}

type AttributeHolder struct {
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	Service       string         `json:""service"`
	Status        string         `json:"status"`
	Timestamp     time.Time      `json:"timestamp"`
	Host          string         `json:"host"`
	SubAttributes []SubAttribute `json:"attributes"`
}

type SubAttribute struct {
	Msg string `json:"msg"`
}
