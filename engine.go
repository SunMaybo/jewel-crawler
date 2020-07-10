package jewel_crawler

import (
	"context"
	"encoding/json"
	"github.com/SunMaybo/jewel-crawler/crawler"
	"github.com/SunMaybo/jewel-crawler/limit"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type CrawlerEngine struct {
	redis    *redis.Client
	limit    *limit.ConcurrentLimit
	Pipeline *crawler.PipeLine
}

func New(redis *redis.Client, concurrents int) *CrawlerEngine {
	return &CrawlerEngine{
		redis:    redis,
		limit:    limit.NewConcurrentLimit(concurrents),
		Pipeline: crawler.New(temp.NewTempStorage(redis)),
	}
}

//开启
func (p *CrawlerEngine) Start(ctx context.Context, channel string, maxExecuteCount int) {

	for {
		t := task.Task{}
		message := <-p.redis.Subscribe(ctx, channel).Channel()
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
				err := p.Push(ctx, channel, task)
				if err != nil {
					zap.S().Fatal(err)
				}
			}
		})

	}

}
func (p *CrawlerEngine) Push(ctx context.Context, channel string, task task.Task) error {
	taskStr, _ := json.Marshal(task)
	return p.redis.Publish(ctx, channel, taskStr).Err()
}
func (p *CrawlerEngine) Close() error {
	return p.redis.Close()
}
