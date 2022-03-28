package wallhaven_sdk_go

import (
	"context"
	"io"
	"net/http"

	"github.com/mitchellh/mapstructure"
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

	return &Client{
		opt: opt,
		hc:  &http.Client{},
	}
}

type Client struct {
	opt *option
	hc  *http.Client
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

func WithBody(body io.Reader) SetDoOpt {
	return func(opt *doOpt) {
		opt.body = body
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
	body   io.Reader

	query  map[string]string
	header map[string]string

	parser Parser
}

func (c *Client) Do(ctx context.Context, doOptFs ...SetDoOpt) error {
	doOpt := newDoOpt()
	for _, f := range doOptFs {
		f(doOpt)
	}

	req, err := http.NewRequestWithContext(ctx, doOpt.method, doOpt.url, doOpt.body)
	if err != nil {
		return err
	}

	query := req.URL.Query()
	for key, value := range doOpt.query {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	for key, value := range doOpt.header {
		req.Header.Add(key, value)
	}

	rsp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	if doOpt.parser == nil {
		return nil
	}
	return doOpt.parser.Parse(req)
}

func sToStruct() {
	mapstructure.Decode(nil, nil)
}
