package wallhaven_sdk_go

import (
	"context"
	"errors"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL = "https://wallhaven.cc/api"
	version = "/v1"
)

var (
	ErrUnknown = errors.New("unknown err")
)

func WithDebug(debug bool) SetOption {
	return func(o *option) {
		o.debug = debug
	}
}

func WithAPIKey(apiKey string) SetOption {
	return func(o *option) {
		o.apiKey = apiKey
	}
}

type SetOption func(*option)

type option struct {
	debug  bool
	apiKey string
}

func NewClient(optFs ...SetOption) *Client {
	opt := new(option)
	for _, f := range optFs {
		f(opt)
	}

	c := &Client{
		opt: opt,
		rc:  resty.New(),
	}

	if c.opt.debug {
		c.rc.
			EnableTrace().
			OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
				log.Printf("url: %s\n", request.URL)
				return nil
			}).
			OnAfterResponse(func(client *resty.Client, r *resty.Response) error {
				log.Printf("raw url: %s\n", r.Request.RawRequest.URL.RequestURI())
				log.Printf("body: %s\n", r.Body())
				return nil
			})
	}
	return c
}

type Client struct {
	opt *option
	rc  *resty.Client
}

func (c *Client) r(ctx context.Context) *resty.Request {
	r := c.rc.R()
	r.SetContext(ctx)
	r.SetHeader("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")

	if c.opt.apiKey != "" {
		r.SetQueryParam("apikey", c.opt.apiKey)
	}
	if c.opt.debug {
		r.EnableTrace()
	}
	return r
}
