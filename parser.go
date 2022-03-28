package wallhaven_sdk_go

import (
	"net/http"
)

type Parser interface {
	Parse(*http.Request) error
}

type jsonParser struct {
}
