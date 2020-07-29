package model

// Resp :
type Resp struct {
	Method string  `json:"method,string"`
	Result float64 `json:"result,float"`
}

// Req :
type Req struct {
	A float64 `json:"a,float"`
	B float64 `json:"b,float"`
}
