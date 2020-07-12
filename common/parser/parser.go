package parser

import (
	"errors"
	"github.com/storyicon/graphquery"
	"go.uber.org/zap"
)

func Parser(content, pattern string) (interface{}, error) {
	response := graphquery.ParseFromString(content, pattern)
	if response.Data != nil {
		return response.Data, nil
	} else {
		zap.S().Warnf("parser error ", "response", response)
		return nil, errors.New("parser error data is nil")
	}
}
