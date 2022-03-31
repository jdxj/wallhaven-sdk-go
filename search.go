package wallhaven_sdk_go

import (
	"fmt"
)

type Color string

const (
	Color660000 Color = "660000"
	Color990000       = "990000"
	Colorcc0000       = "cc0000"
	Colorcc3333       = "cc3333"
	Colorea4c88       = "ea4c88"
	Color993399       = "993399"
	Color663399       = "663399"
	Color333399       = "333399"
	Color0066cc       = "0066cc"
	Color0099cc       = "0099cc"
	Color66cccc       = "66cccc"
	Color77cc33       = "77cc33"
	Color669900       = "669900"
	Color336600       = "336600"
	Color666600       = "666600"
	Color999900       = "999900"
	Colorcccc33       = "cccc33"
	Colorffff00       = "ffff00"
	Colorffcc33       = "ffcc33"
	Colorff9900       = "ff9900"
	Colorff6600       = "ff6600"
	Colorcc6633       = "cc6633"
	Color996633       = "996633"
	Color663300       = "663300"
	Color000000       = "000000"
	Color999999       = "999999"
	Colorcccccc       = "cccccc"
	Colorffffff       = "ffffff"
	Color424153       = "424153"
)

type Category int

const (
	People Category = 1 << iota
	Anime
	General
)

func (c Category) String() (s string) {
	for i := 2; i >= 0; i-- {
		s += fmt.Sprintf("%d", (c>>i)&1)
	}
	return
}

type Purity int

const ()

const (
	PNG = "png"
	JPG = "jpg"
)

type Query struct {
	param string
}

func (q Query) String() string {
	return ""
}

func (q Query) FuzzilyTags(tags ...string) {
}

func (q Query) ExcludeTags(tags ...string) {}

func (q Query) MustTags(tags ...string) {

}

func (q Query) ExactTags(tags ...string) {

}

func (q Query) Type(t string) {

}

func (q Query) Like(id string) {

}

type SearchReq struct {
	Query Query
	Cate  Category
}
