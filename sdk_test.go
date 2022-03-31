package wallhaven_sdk_go

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var (
	client *Client
)

func TestMain(t *testing.M) {
	client = NewClient(WithDebug(true))
	os.Exit(t.Run())
}

func TestClient_GetWallpaper(t *testing.T) {
	req := &GetWallpaperReq{
		ID: "k7v9yq",
	}

	rsp, err := client.GetWallpaper(context.Background(), req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

type Person struct {
	Age int `mapstructure:"age"`
}

func TestStructToMap(t *testing.T) {
	p := &Person{
		Age: 10,
	}
	m := structToMap(p)
	fmt.Printf("%v\n", m)
}

func TestCategory_String(t *testing.T) {
	cases := []struct {
		name string
		cate Category
		str  string
	}{
		{
			name: "case1",
			cate: People,
			str:  "001",
		},
		{
			name: "case2",
			cate: Anime,
			str:  "010",
		},
		{
			name: "case3",
			cate: General,
			str:  "100",
		},
		{
			name: "case4",
			cate: People | Anime,
			str:  "011",
		},
		{
			name: "case5",
			cate: People | General,
			str:  "101",
		},
		{
			name: "case6",
			cate: Anime | General,
			str:  "110",
		},
		{
			name: "case7",
			cate: People | Anime | General,
			str:  "111",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.cate.String() != c.str {
				t.Errorf("%+v", c)
			}
		})
	}
}
