package jewel_crawler

import (
	"context"
	"github.com/SunMaybo/jewel-crawler/common"
	"github.com/SunMaybo/jewel-crawler/crawler"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestCrawlerEngine_Start(t *testing.T) {
	engine := New(&Config{
		Redis: &redis.Options{
			Addr: "127.0.0.1:6379",
		},
		Queue:      "shtml",
		Concurrent: 10,
	})
	SetLogLevel("info")
	id := common.GenerateRandomID()
	err := engine.push(context.Background(), "shtml", task.Task{
		CrawlerName: "default_html_crawler",
		CrawlerUrl:  "https://juejin.im/post/5f0ad0db6fb9a07e777e9e81",
		Website:     "juejin.im",
		Method:      "GET",
		Depth:       1,
		Index:       1,
		Timeout:     5 * time.Second,
		GlobalId:    id,
		ParentId:    id,
		TaskId:      id,
		Time:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatal(err)
	}
	engine.Pipeline.AddCrawler("default_html_crawler", &crawler.DefaultHtmlCrawler{})
	engine.Start(context.Background(), 3)
}
