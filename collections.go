package wallhaven_sdk_go

import (
	"context"
	"fmt"
	"strconv"
)

type GetCollectionsReq struct {
	Username string
}

func (gcr *GetCollectionsReq) API() string {
	path := "/collections"
	if gcr.Username != "" {
		path += "/{username}"
	}
	return baseURL + version + path
}

func (gcr *GetCollectionsReq) Map() map[string]string {
	return map[string]string{
		"username": gcr.Username,
	}
}

type Collection struct {
	Id     int    `json:"id"`
	Label  string `json:"label"`
	Views  int    `json:"views"`
	Public int    `json:"public"`
	Count  int    `json:"count"`
}

type GetCollectionsRsp struct {
	Collections []Collection `json:"data"`
}

func (c *Client) GetCollections(ctx context.Context, req *GetCollectionsReq) (*GetCollectionsRsp, error) {
	rsp, err := c.r(ctx).
		SetPathParams(req.Map()).
		SetResult(&GetCollectionsRsp{}).
		Get(req.API())
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*GetCollectionsRsp), nil
}

type GetCollectionWallpapersReq struct {
	Username string
	ID       int
}

func (gcw *GetCollectionWallpapersReq) API() string {
	path := "/collections/{username}/{id}"
	return baseURL + version + path
}

func (gcw *GetCollectionWallpapersReq) Map() map[string]string {
	return map[string]string{
		"username": gcw.Username,
		"id":       strconv.Itoa(gcw.ID),
	}
}

type GetCollectionWallpapersRsp struct {
	Wallpapers []Wallpaper `json:"data"`
	Meta       Meta        `json:"meta"`
}

func (c *Client) GetCollectionWallpapers(ctx context.Context, req *GetCollectionWallpapersReq) (
	*GetCollectionWallpapersRsp, error) {
	rsp, err := c.r(ctx).
		SetPathParams(req.Map()).
		SetResult(&GetCollectionWallpapersRsp{}).
		Get(req.API())
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*GetCollectionWallpapersRsp), nil
}
