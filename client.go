package wallhaven_sdk_go

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	baseURL = "https://wallhaven.cc/api"
	version = "/v1"
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
				fmt.Printf("body: %s\n", r.Body())
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

func WithMethod(method string) SetDoOpt {
	return func(opt *doOpt) {
		opt.method = method
	}
}

func WithURL(url string) SetDoOpt {
	return func(opt *doOpt) {
		opt.url = url
	}
}

// todo: set file, multipart file
func WithBody(body io.Reader) SetDoOpt {
	return func(opt *doOpt) {
		//opt.body = body
	}
}

func WithQuery(query map[string]string) SetDoOpt {
	return func(opt *doOpt) {
		opt.query = query
	}
}

func WithHeader(header map[string]string) SetDoOpt {
	return func(opt *doOpt) {
		opt.header = header
	}
}

func WithParser(parser Parser) SetDoOpt {
	return func(opt *doOpt) {
		opt.parser = parser
	}
}

type SetDoOpt func(*doOpt)

func newDoOpt() *doOpt {
	do := &doOpt{
		method: http.MethodGet,
	}
	return do
}

type doOpt struct {
	method string
	url    string

	query  map[string]string
	header map[string]string

	parser Parser
}
