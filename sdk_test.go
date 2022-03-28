package wallhaven_sdk_go

import (
	"context"
	"net/http"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func TestNewReq(t *testing.T) {
	rsp, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	if rsp.URL != nil {
		t.Logf("url ok")
	}
}

type Person struct {
	Name        string `mapstructure:"name"`
	FirstSecond string `mapstructure:"f_s"`
}

func TestMapStruct(t *testing.T) {
	p := Person{
		Name:        "abc",
		FirstSecond: "def",
	}
	var c map[string]string
	err := mapstructure.Decode(p, &c)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("%v\n", c)
}
