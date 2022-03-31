package wallhaven_sdk_go

import (
	"context"
)

type GetWallpaperInfoReq struct {
	ID string `mapstructure:"id"`
}

type Uploader struct {
	Username string            `json:"username"`
	Group    string            `json:"group"`
	Avatar   map[string]string `json:"avatar"`
}

type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
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

type GetWallpaperInfoRsp struct {
	Id         string   `json:"id"`
	Url        string   `json:"url"`
	ShortUrl   string   `json:"short_url"`
	Uploader   Uploader `json:"uploader"`
	Views      int      `json:"views"`
	Favorites  int      `json:"favorites"`
	Source     string   `json:"source"`
	Purity     string   `json:"purity"`
	Category   string   `json:"category"`
	DimensionX int      `json:"dimension_x"`
	DimensionY int      `json:"dimension_y"`
	Resolution string   `json:"resolution"`
	Ratio      string   `json:"ratio"`
	FileSize   int      `json:"file_size"`
	FileType   string   `json:"file_type"`
	CreatedAt  string   `json:"created_at"`
	Colors     []string `json:"colors"`
	Path       string   `json:"path"`
	Thumbs     Thumbs   `json:"thumbs"`
	Tags       []Tag    `json:"tags"`
}

func (c *Client) GetWallpaperInfo(ctx context.Context, req *GetWallpaperInfoReq) (*GetWallpaperInfoRsp, error) {
	var (
		url  = baseURL + version + "/w/{id}"
		wrap = newRsp(&GetWallpaperInfoRsp{})
	)
	rsp, err := c.r(ctx).
		SetPathParams(structToMap(req)).
		SetResult(wrap).
		Get(url)
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*response).Data.(*GetWallpaperInfoRsp), nil
}
