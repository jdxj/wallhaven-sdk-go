package wallhaven_sdk_go

func newRsp(data interface{}) *response {
	return &response{
		Data: data,
	}
}

type response struct {
	Data interface{} `json:"data"`
}

type MetaQuery struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type Meta struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
	Query       interface{} `json:"query"`
	Seed        interface{} `json:"seed"`
}
