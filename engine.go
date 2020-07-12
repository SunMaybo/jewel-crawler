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
	"time"
)

type CrawlerEngine struct {
	redis    *redis.Client
	limit    *limit.ConcurrentLimit
	Pipeline *crawler.PipeLine
	queue    string
}

func SetLogLevel(level string) {
	logs.GetLog(level)
}

type Config struct {
	Addr       string
	Password   string
	Db         int
	Queue      string
	Concurrent int
}

func New(cfg *Config) *CrawlerEngine {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return &CrawlerEngine{
		redis:    rdb,
		queue:    cfg.Queue,
		limit:    limit.NewConcurrentLimit(cfg.Concurrent),
		Pipeline: crawler.New(cfg.Queue, temp.NewTempStorage(rdb)),
	}
}

//开启
func (p *CrawlerEngine) Start(ctx context.Context, maxExecuteCount int) {
	for {
		result, err := p.redis.LPop(ctx, p.queue).Result()
		if err != nil && err != redis.Nil {
			panic(err)
		}
		if err != nil && redis.Nil == err {
			time.Sleep(5 * time.Millisecond)
			continue
		}
		t := task.Task{}
		err = json.Unmarshal([]byte(result), &t)
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
				err := p.push(ctx, p.queue, task)
				if err != nil {
					logs.S.Fatal(err)
				}
			}
		})
	}

}
func (p *CrawlerEngine) push(ctx context.Context, queue string, task task.Task) error {
	taskStr, _ := json.Marshal(task)
	return p.redis.LPush(ctx, queue, taskStr).Err()
}
func (p *CrawlerEngine) Close() error {
	return p.redis.Close()
}
