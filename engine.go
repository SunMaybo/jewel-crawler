package jewel_crawler

import (
	"context"
	"encoding/json"
	"github.com/SunMaybo/jewel-crawler/crawler"
	"github.com/SunMaybo/jewel-crawler/limit"
	logs "github.com/SunMaybo/jewel-crawler/logs"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type CrawlerEngine struct {
	redis    *redis.Client
	limit    *limit.ConcurrentLimit
	Pipeline *crawler.PipeLine
	channel  string
}

func SetLogLevel(level string) {
	logs.GetLog(level)
}

func New(redis *redis.Client, channel string, concurrent int) *CrawlerEngine {
	return &CrawlerEngine{
		redis:    redis,
		limit:    limit.NewConcurrentLimit(concurrent),
		Pipeline: crawler.New(channel, temp.NewTempStorage(redis)),
	}
}

//开启
func (p *CrawlerEngine) Start(ctx context.Context, maxExecuteCount int) {
	for {
		t := task.Task{}
		message := <-p.redis.Subscribe(ctx, p.channel).Channel()
		err := json.Unmarshal([]byte(message.Payload), &t)
		if err != nil {
			panic(err)
		}
		p.limit.Acquire(t, func(task task.Task) {
			defer p.limit.Free()
			//todo
			if task.Retry >= maxExecuteCount {
				return
			}
			err := p.Pipeline.Invoke(ctx, task)
			if err != nil {
				task.Retry += 1
				err := p.push(ctx, p.channel, task)
				if err != nil {
					zap.S().Fatal(err)
				}
			}
		})

	}

}
func (p *CrawlerEngine) push(ctx context.Context, channel string, task task.Task) error {
	taskStr, _ := json.Marshal(task)
	return p.redis.Publish(ctx, channel, taskStr).Err()
}
func (p *CrawlerEngine) Close() error {
	return p.redis.Close()
}
