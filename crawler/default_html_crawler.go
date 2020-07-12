package crawler

import (
	"github.com/SunMaybo/go-readability"
	"github.com/SunMaybo/jewel-crawler/common/spider"
	"github.com/SunMaybo/jewel-crawler/logs"
	"strings"
	"time"
)

type DefaultHtmlCrawler struct {
}

func (dhc *DefaultHtmlCrawler) Collect(event CollectEvent) (string, error) {
	s := spider.NewShtmlSpider(1 * 1024 * 1024)
	resp, err := s.Do(spider.Request{
		Url:     event.Task.CrawlerUrl,
		Method:  event.Task.Method,
		Param:   event.Task.Param,
		Headers: event.Task.Header,
		Timeout: event.Task.Timeout,
	})
	if err != nil {
		return "", err
	}
	return resp.GetContent(), nil
}
func (dhc *DefaultHtmlCrawler) Parser(event ParserEvent) (map[string]interface{}, error) {
	data, err := parserArticleWithReadability(event.Content, event.Task.CrawlerUrl)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dhc *DefaultHtmlCrawler) Storage(event StorageEvent) error {
	logs.S.Info(event.Data)
	return nil
}

func parserArticleWithReadability(html string, url string) (map[string]interface{}, error) {
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
