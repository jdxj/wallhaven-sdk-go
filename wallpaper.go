package wallhaven_sdk_go

import (
	"context"
	"errors"
)

type GetWallpaperReq struct {
	ID string
}

func (gwr *GetWallpaperReq) API() string {
	return baseURL + version + "/w/{id}"
}

func (gwr *GetWallpaperReq) Map() map[string]string {
	return map[string]string{
		"id": gwr.ID,
	}
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

type Wallpaper struct {
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

func (c *Client) GetWallpaper(ctx context.Context, req *GetWallpaperReq) (*Wallpaper, error) {
	wrap := newRsp(&Wallpaper{})
	rsp, err := c.r(ctx).
		SetPathParams(req.Map()).
		SetResult(wrap).
		Get(req.API())
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, errors.New(rsp.Status())
	}

	return rsp.Result().(*response).Data.(*Wallpaper), nil
}
