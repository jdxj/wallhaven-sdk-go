package wallhaven_sdk_go

import (
	"context"
)

type GetTagInfoReq struct {
	ID int `mapstructure:"id"`
}

type GetTagInfoRsp struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryId int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

func (c *Client) GetTagInfo(ctx context.Context, req *GetTagInfoReq) (*GetTagInfoRsp, error) {
	var (
		url  = baseURL + version + "/tag/{id}"
		wrap = newRsp(&GetTagInfoRsp{})
	)
	rsp, err := c.r(ctx).
		SetPathParams(structToMap(req)).
		SetResult(wrap).
		Get(url)
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*response).Data.(*GetTagInfoRsp), nil
}
