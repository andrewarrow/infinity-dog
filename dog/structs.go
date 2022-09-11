package dog

type LogResponse struct {
	Data []AttributeHolder `json:"data"`
}

type AttributeHolder struct {
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	Service       string         `json:""service"`
	Status        string         `json:"status"`
	Timestamp     string         `json:"timestamp"`
	Host          string         `json:"host"`
	SubAttributes []SubAttribute `json:"attributes"`
}

type SubAttribute struct {
	Msg string `json:"msg"`
}
