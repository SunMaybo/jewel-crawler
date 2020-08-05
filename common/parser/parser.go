package parser

import (
	"errors"
	"github.com/SunMaybo/go-readability"
	"github.com/SunMaybo/jewel-crawler/logs"
	"github.com/storyicon/graphquery"
	"strings"
	"time"
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

func ParserArticleWithReadability(html string, url string) (map[string]interface{}, error) {
	starTime := time.Now()
	parser := readability.NewParser()
	article, err := parser.Parse(strings.NewReader(html), url)
	logs.S.Infow("解析耗时", "interval", time.Since(starTime).String())
	if err != nil {
		logs.S.Warn(err)
		logs.S.Warnw("规则匹配耗时", "interval", time.Since(starTime).String())
		return nil, err
	} else {
		resultMap := make(map[string]interface{})
		resultMap["title"] = article.Title
		resultMap["author"] = article.Byline
		resultMap["source_name"] = article.SourceName
		resultMap["content"] = article.Content
		logs.S.Infow("规则匹配耗时", "interval", time.Since(starTime).String())
		return resultMap, nil
	}

}
