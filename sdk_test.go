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
