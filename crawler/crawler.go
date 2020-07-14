package crawler

import (
	"github.com/SunMaybo/jewel-crawler/common"
	"github.com/SunMaybo/jewel-crawler/common/parser"
	"github.com/SunMaybo/jewel-crawler/common/spider"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
)

type Event struct {
}

type CollectEvent struct {
	Event
	Task  task.Task
	Queue string
	Temp  temp.Temp
}
type ParserEvent struct {
	Event
	Task    task.Task
	Queue   string
	Temp    temp.Temp
	Content string
}
type StorageEvent struct {
	Event
	Task        task.Task
	Queue       string
	Data        map[string]interface{}
	TempStorage *temp.TempStorage
}

func (event *Event) ApiSpider(size int) spider.Spider {
	return spider.NewApiSpider(size)
}
func (event *Event) ShtmlSpider(size int) spider.Spider {
	return spider.NewShtmlSpider(size)
}
func (event *Event) DhtmlSpider() spider.Spider {
	return spider.NewDhtmlSpider()
}
func (event *Event) FileSpider(size int) spider.Spider {
	return spider.NewFileSpider(size)
}
func (event *Event) Parser(content, pattern string) (interface{}, error) {
	return parser.Parser(content, pattern)
}
func (event *Event) ReadabilityParser(html, url string) (map[string]interface{}, error) {
	return parser.ParserArticleWithReadability(html, url)
}
func (event *Event) Signature(obj interface{}) string {
	return common.Signature(obj)
}

func (event *Event) GenerateRandomID() string {
	return common.GenerateRandomID()
}

func (event *Event) SignatureMap(data map[string]string) string {
	return common.SignatureMap(data)
}

type Crawler interface {
	Collect(event CollectEvent) (string, error)
	Parser(event ParserEvent) (map[string]interface{}, error)
	Storage(event StorageEvent) error
}
