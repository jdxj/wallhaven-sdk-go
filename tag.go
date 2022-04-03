package wallhaven_sdk_go

import (
	"context"
	"strconv"
)

type GetTagReq struct {
	ID int
}

func (gti *GetTagReq) API() string {
	return baseURL + version + "/tag/{id}"
}

func (gti *GetTagReq) Map() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(gti.ID),
	}
}

type Tag struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryId int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	CreatedAt  string `json:"created_at"`
}

func (c *Client) GetTag(ctx context.Context, req *GetTagReq) (*Tag, error) {
	wrap := newRsp(&Tag{})
	rsp, err := c.r(ctx).
		SetPathParams(req.Map()).
		SetResult(wrap).
		Get(req.API())
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*response).Data.(*Tag), nil
}
