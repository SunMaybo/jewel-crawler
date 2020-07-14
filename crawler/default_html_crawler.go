package crawler

import (
	"github.com/SunMaybo/jewel-crawler/common/spider"
	"github.com/SunMaybo/jewel-crawler/logs"
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
	data, err := event.ReadabilityParser(event.Content, event.Task.CrawlerUrl)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dhc *DefaultHtmlCrawler) Storage(event StorageEvent) error {
	logs.S.Info(event.Data)
	return nil
}
