package parser

import (
	"errors"
	"github.com/SunMaybo/jewel-crawler/logs"
	"github.com/storyicon/graphquery"
)

func Parser(content, pattern string) (interface{}, error) {
	response := graphquery.ParseFromString(content, pattern)
	if response.Data != nil {
		return response.Data, nil
	} else {
		logs.S.Warnw("parser error ", "response", response)
		return nil, errors.New("parser error data is nil")
	}
}
