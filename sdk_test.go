package wallhaven_sdk_go

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
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

func TestClient_GetTagInfo(t *testing.T) {
	req := &GetTagReq{
		ID: 372,
	}
	rsp, err := client.GetTag(context.Background(), req)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetCollections(t *testing.T) {
	rsp, err := client.GetCollections(context.Background(), &GetCollectionsReq{
		Username: "jdxj",
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.Collections {
		fmt.Printf("%+v\n", v)
	}
}

func TestClient_GetCollectionWallpapers(t *testing.T) {
	rsp, err := client.GetCollectionWallpapers(context.Background(), &GetCollectionWallpapersReq{
		Username: "jdxj",
		ID:       1151634,
	})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestClient_GetSettings(t *testing.T) {
	rsp, err := client.GetSettings(context.Background())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", rsp)
}

func TestResolutions_Map(t *testing.T) {
	r := &Resolutions{}
	r.SetAtLeast(R_1280x720).
		SetExact(R_1600x1000).
		SetCustom(123, 456)
	res := r.Map()
	fmt.Printf("%v\n", res)
}

func TestJoin(t *testing.T) {
	ss := []string{""}
	res := strings.Join(ss, "-")
	fmt.Printf("res: %s\n", res)
}

func TestQuery_String(t *testing.T) {
	q := &Query{
		fuzzy:    []string{},
		exclude:  []string{"people", "apple"},
		must:     []string{},
		exact:    123,
		username: "king",
		typ:      PNG,
		like:     "fff",
	}
	fmt.Printf("%s\n", q)
}

func TestConvert(t *testing.T) {
	var n rune = 48
	n = n + 2
	fmt.Printf("%s", string(n))
}

func TestCategory_String2(t *testing.T) {
	cate := People | Anime
	fmt.Printf("%s\n", cate)
}

func TestString(t *testing.T) {
	ss := "abc"
	for _, v := range ss {
		fmt.Printf("%d\n", v)
		t := reflect.TypeOf(v)
		fmt.Printf("%s\n", t)
	}

	rs := []int32{48, 49}
	fmt.Printf("%s\n", string(rs))
}

func TestClient_Search(t *testing.T) {
	r := &SearchReq{
		Query: Query{
			fuzzy:    []string{},
			exclude:  []string{},
			must:     []string{"apple", "people"},
			exact:    0,
			username: "",
			typ:      "",
			like:     "",
		},
	}
	fmt.Printf("req: %v\n", r.Map())
	rsp, err := client.Search(context.Background(), r)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range rsp.Wallpapers {
		fmt.Printf("%+v\n", v)
	}
	fmt.Printf("meta: %+v\n", rsp.Meta)
}
