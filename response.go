package wallhaven_sdk_go

func newRsp(data interface{}) *response {
	return &response{
		Data: data,
	}
}

type response struct {
	Data interface{} `json:"data"`
}
