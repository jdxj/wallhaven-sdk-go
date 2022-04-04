package wallhaven_sdk_go

import (
	"context"
	"fmt"
)

type GetSettingsRsp struct {
	ThumbSize     string   `json:"thumb_size"`
	PerPage       string   `json:"per_page"`
	Purity        []string `json:"purity"`
	Categories    []string `json:"categories"`
	Resolutions   []string `json:"resolutions"`
	AspectRatios  []string `json:"aspect_ratios"`
	TopListRange  string   `json:"toplist_range"`
	TagBlacklist  []string `json:"tag_blacklist"`
	UserBlacklist []string `json:"user_blacklist"`
}

func (c *Client) GetSettings(ctx context.Context) (*GetSettingsRsp, error) {
	wrap := newRsp(&GetSettingsRsp{})
	api := baseURL + version + "/settings"
	rsp, err := c.r(ctx).
		SetResult(wrap).
		Get(api)
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*response).Data.(*GetSettingsRsp), nil
}
