package wallhaven_sdk_go

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

func structToMap(s interface{}) map[string]string {
	var m map[string]string
	err := mapstructure.Decode(s, &m)
	if err != nil {
		log.Printf("structToMap: %s", err)
	}
	return m
}
