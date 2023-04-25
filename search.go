package wallhaven_sdk_go

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type Color = string

const (
	C_660000 Color = "660000"
	C_990000 Color = "990000"
	C_cc0000 Color = "cc0000"
	C_cc3333 Color = "cc3333"
	C_ea4c88 Color = "ea4c88"
	C_993399 Color = "993399"
	C_663399 Color = "663399"
	C_333399 Color = "333399"
	C_0066cc Color = "0066cc"
	C_0099cc Color = "0099cc"
	C_66cccc Color = "66cccc"
	C_77cc33 Color = "77cc33"
	C_669900 Color = "669900"
	C_336600 Color = "336600"
	C_666600 Color = "666600"
	C_999900 Color = "999900"
	C_cccc33 Color = "cccc33"
	C_ffff00 Color = "ffff00"
	C_ffcc33 Color = "ffcc33"
	C_ff9900 Color = "ff9900"
	C_ff6600 Color = "ff6600"
	C_cc6633 Color = "cc6633"
	C_996633 Color = "996633"
	C_663300 Color = "663300"
	C_000000 Color = "000000"
	C_999999 Color = "999999"
	C_cccccc Color = "cccccc"
	C_ffffff Color = "ffffff"
	C_424153 Color = "424153"
)

type Category int32

const (
	People Category = 1 << iota
	Anime
	General
)

func (c Category) String() string {
	var str []int32
	for i := 2; i >= 0; i-- {
		str = append(str, int32(c>>i)&1+48)
	}

	res := string(str)
	if res == "000" {
		return ""
	}
	return res
}

type AIArt bool

func (a AIArt) String() string {
	if a {
		return "0"
	} else {
		return "1"
	}
}

type Purity int32

const (
	NSFW Purity = 1 << iota
	Sketchy
	SFW
)

func (p Purity) String() string {
	var str []int32
	for i := 2; i >= 0; i-- {
		str = append(str, int32(p>>i)&1+48)
	}

	res := string(str)
	if res == "000" {
		return ""
	}
	return res
}

type Sorting = string

const (
	DateAdded Sorting = "date_added"
	Relevance Sorting = "relevance"
	Random    Sorting = "random"
	Views     Sorting = "views"
	Favorites Sorting = "favorites"
	TopList   Sorting = "toplist"
	Hot       Sorting = "hot"
)

type Order = string

const (
	Desc Order = "desc"
	Asc  Order = "asc"
)

type TopRange = string

const (
	D1 TopRange = "1d"
	D3 TopRange = "3d"
	W1 TopRange = "1w"
	M1 TopRange = "1M"
	M3 TopRange = "3M"
	M6 TopRange = "6M"
	Y1 TopRange = "1y"
)

type Resolution = string

const (
	// Ultrawide
	R_2560x1080 Resolution = "2560x1080"
	R_3440x1440 Resolution = "3440x1440"
	R_3840x1600 Resolution = "3840x1600"

	// 16:9
	R_1280x720  = "1280x720"
	R_1600x900  = "1600x900"
	R_1920x1080 = "1920x1080"
	R_2560x1440 = "2560x1440"
	R_3840x2160 = "3840x2160"

	// 16:10
	R_1280x800  = "1280x800"
	R_1600x1000 = "1600x1000"
	R_1920x1200 = "1920x1200"
	R_2560x1600 = "2560x1600"
	R_3840x2400 = "3840x2400"

	// 4:3
	R_1280x960  = "1280x960"
	R_1600x1200 = "1600x1200"
	R_1920x1440 = "1920x1440"
	R_2560x1920 = "2560x1920"
	R_3840x2880 = "3840x2880"

	// 5:4
	R_1280x1024 = "1280x1024"
	R_1600x1280 = "1600x1280"
	R_1920x1536 = "1920x1536"
	R_2560x2048 = "2560x2048"
	R_3840x3072 = "3840x3072"
)

type Resolutions struct {
	atLeast Resolution
	exact   []Resolution
	custom  Resolution
}

func (r *Resolutions) Map() map[string]string {
	var (
		m              = make(map[string]string)
		KeyAtLeast     = "atleast"
		KeyResolutions = "resolutions"
	)
	if r.atLeast != "" {
		if r.custom != "" {
			r.atLeast = r.custom
		}
		m[KeyAtLeast] = r.atLeast
	} else {
		if r.custom != "" {
			r.exact = append(r.exact, r.custom)
		}
		m[KeyResolutions] = strings.Join(r.exact, ",")
	}
	return m
}

func (r *Resolutions) SetAtLeast(al Resolution) *Resolutions {
	r.atLeast = al
	r.exact = r.exact[:0]
	return r
}

func (r *Resolutions) SetExact(rs ...Resolution) *Resolutions {
	r.exact = append(r.exact, rs...)
	r.atLeast = ""
	return r
}

func (r *Resolutions) SetCustom(width, height int) *Resolutions {
	r.custom = fmt.Sprintf("%dx%d", width, height)
	return r
}

type Ratio = string

const (
	Landscape Ratio = "landscape"
	Portrait  Ratio = "portrait"

	// Wide
	O_16x9  Ratio = "16x9"
	O_16x10 Ratio = "16x10"

	// Ultrawide
	O_21x9 Ratio = "21x9"
	O_32x9 Ratio = "32x9"
	O_48x9 Ratio = "48x9"

	// Portrait
	O_9x16  Ratio = "9x16"
	O_10x16 Ratio = "10x16"
	O_9x18  Ratio = "9x18"

	// Square
	O_1x1 Ratio = "1x1"
	O_3x2 Ratio = "3x2"
	O_4x3 Ratio = "4x3"
	O_5x4 Ratio = "5x4"
)

type Ratios struct {
	ratios []Ratio
}

func (r *Ratios) String() string {
	return strings.Join(r.ratios, ",")
}

func (r *Ratios) AddRatio(ratios ...Ratio) *Ratios {
	r.ratios = append(r.ratios, ratios...)
	return r
}

type Type = string

const (
	PNG Type = "png"
	JPG Type = "jpg"
)

type Query struct {
	// search fuzzily for a tag/keyword
	fuzzy []string
	// exclude a tag/keyword
	exclude []string
	// must have tag
	must []string
	// user uploads
	username string
	// exact tag search (can not be combined)
	exact string
	// search for file type (jpg = jpeg)
	typ Type
	// find wallpapers with similar tags
	like string
}

func (q *Query) AddFuzzy(tag string) {
	q.fuzzy = append(q.fuzzy, TrimSpaceAndSplit(tag)...)
}

func (q *Query) AddExclude(tag string) {
	for _, v := range TrimSpaceAndSplit(tag) {
		q.exclude = append(q.exclude, "-"+v)
	}
}

func (q *Query) AddMust(tag string) {
	for _, v := range TrimSpaceAndSplit(tag) {
		q.must = append(q.must, "+"+v)
	}
}

func (q *Query) SetUsername(username string) {
	q.username = "@" + strings.TrimSpace(username)
}

func (q *Query) SetExact(tag int) {
	q.exact = "id:" + strconv.Itoa(tag)
}

func (q *Query) SetType(typ Type) {
	q.typ = "type:" + strings.TrimSpace(typ)
}

func (q *Query) SetLike(like string) {
	q.like = "like:" + strings.TrimSpace(like)
}

func (q *Query) String() (result string) {
	var s []string
	if len(q.fuzzy) != 0 {
		s = append(s, strings.Join(q.fuzzy, " "))
	}
	if len(q.exclude) != 0 {
		s = append(s, strings.Join(q.exclude, " "))
	}
	if len(q.must) != 0 {
		s = append(s, strings.Join(q.must, " "))
	}
	if q.username != "" {
		s = append(s, q.username)
	}
	if q.exact != "" {
		s = append(s, q.exact)
	}
	if q.typ != "" {
		s = append(s, q.typ)
	}
	if q.like != "" {
		s = append(s, q.like)
	}

	if len(s) != 0 {
		result = strings.Join(s, " ")
	}
	return
}

type SearchReq struct {
	Query
	Category
	AIArt
	Purity
	Sorting
	Order
	TopRange
	Resolutions
	Ratios
	Color

	Page int
	Seed string
}

func (sr *SearchReq) API() string {
	return baseURL + version + "/search"
}

func (sr *SearchReq) Map() map[string]string {
	m := make(map[string]string)

	m["q"] = sr.Query.String()
	m["categories"] = sr.Category.String()
	m["ai_art_filter"] = sr.AIArt.String()
	m["purity"] = sr.Purity.String()
	m["sorting"] = sr.Sorting
	m["order"] = sr.Order
	m["topRange"] = sr.TopRange
	m["ratios"] = sr.Ratios.String()
	m["colors"] = sr.Color
	m["seed"] = sr.Seed

	for k, v := range sr.Resolutions.Map() {
		m[k] = v
	}
	if sr.Page != 0 {
		m["page"] = strconv.Itoa(sr.Page)
	}

	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}
	return m
}

type SearchRsp struct {
	Wallpapers []Wallpaper `json:"data"`
	Meta       Meta        `json:"meta"`
}

func (c *Client) Search(ctx context.Context, req *SearchReq) (*SearchRsp, error) {
	rsp, err := c.r(ctx).
		SetQueryParams(req.Map()).
		SetResult(&SearchRsp{}).
		Get(req.API())
	if err != nil {
		return nil, err
	}
	if rsp.IsError() {
		return nil, fmt.Errorf("%w: %s", ErrUnknown, rsp.Status())
	}

	return rsp.Result().(*SearchRsp), nil
}
