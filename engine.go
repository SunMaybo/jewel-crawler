package jewel_crawler

import (
	"context"
	"encoding/json"
	"github.com/SunMaybo/jewel-crawler/crawler"
	"github.com/SunMaybo/jewel-crawler/limit"
	logs "github.com/SunMaybo/jewel-crawler/logs"
	"github.com/SunMaybo/jewel-crawler/sync"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
	"github.com/go-redis/redis/v8"
	"github.com/yangwenmai/ratelimit/simpleratelimit"
	"time"
)

type CrawlerEngine struct {
	redis         *redis.Client
	limit         *limit.ConcurrentLimit
	Pipeline      *crawler.PipeLine
	queue         string
	Concurrent    int
	consumerQueue string
	RateLimiter   *simpleratelimit.RateLimiter
	CallBack      func(task task.Task, err error)
}

func SetLogLevel(level string) {
	logs.GetLog(level)
}

type Config struct {
	Redis         *redis.Options
	Queue         string
	ConsumerQueue string
	Concurrent    int
}

func New(cfg *Config) *CrawlerEngine {
	rdb := redis.NewClient(cfg.Redis)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logs.S.Fatal(err)
	}
	if cfg.Concurrent <= 0 {
		cfg.Concurrent = 3
	}
	return &CrawlerEngine{
		redis:         rdb,
		Concurrent:    cfg.Concurrent,
		queue:         cfg.Queue,
		consumerQueue: cfg.ConsumerQueue,
		limit:         limit.NewConcurrentLimit(cfg.Concurrent),
		RateLimiter:   simpleratelimit.New(cfg.Concurrent, time.Second),
		Pipeline:      crawler.New(cfg.Queue, temp.NewTempStorage(rdb)),
	}
}

//开启
func (p *CrawlerEngine) Start(ctx context.Context, maxExecuteCount int) {
	if maxExecuteCount <= 0 {
		maxExecuteCount = 1
	}
	logs.S.Infow("当前处理信息", "Concurrent", p.Concurrent, "queue", p.queue, "maxExecuteCount", maxExecuteCount)

	for {
		logs.S.Debug("loop task......")
		result, err := p.redis.LPop(ctx, p.queue).Result()
		if err != nil && err != redis.Nil {
			logs.S.Error(err)
			time.Sleep(3 * time.Second)
			continue
		}
		logs.S.Debug("loop task ok......")
		if err != nil && redis.Nil == err {
			time.Sleep(15 * time.Second)
			logs.S.Debugw("队列为空", "queue", p.queue)
			continue
		}
		tt := task.Task{}
		err = json.Unmarshal([]byte(result), &tt)
		if err != nil {
			logs.S.Fatal(err)
		}
		tt.Redis = p.redis
		for p.RateLimiter.Limit() {
			time.Sleep(100 * time.Millisecond)
		}
		go func(t task.Task) {
			err := p.Pipeline.Invoke(ctx, t)
			if err != nil {
				if t.Retry <= maxExecuteCount {
					t.Retry += 1
					t.Timeout = t.Timeout + t.Timeout/3
					err := p.Push(ctx, p.queue, t)
					if err != nil {
						logs.S.Warn(err)
					}
					logs.S.Warnw("处理失败，进行重试", "queue", p.queue, "err", err)
				} else {
					if p.CallBack != nil {
						p.CallBack(t, err)
					}
					logs.S.Errorw("处理失败，任务丢弃", "queue", p.queue, "err", err)
				}
			} else {
				if p.CallBack != nil {
					p.CallBack(t, err)
				}
				logs.S.Infow("处理成功", "queue", p.queue)
			}
		}(tt)
	}

}

//开启
func (p *CrawlerEngine) StartThread(ctx context.Context, maxExecuteCount int) {
	if maxExecuteCount <= 0 {
		maxExecuteCount = 1
	}
	for i := 0; i < p.Concurrent; i++ {
		go func() {
			for {
				result, err := p.redis.LPop(ctx, p.queue).Result()
				if err != nil && err != redis.Nil {
					panic(err)
					time.Sleep(3 * time.Second)
					continue
				}
				if err != nil && redis.Nil == err {
					time.Sleep(15 * time.Second)
					logs.S.Debugw("队列为空", "queue", p.queue)
					continue
				}
				t := task.Task{}
				err = json.Unmarshal([]byte(result), &t)
				if err != nil {
					logs.S.Error(err)
					panic(err)
				}
				t.Redis = p.redis
				err = p.Pipeline.Invoke(ctx, t)
				if err != nil {
					if t.Retry <= maxExecuteCount {
						t.Retry += 1
						t.Timeout = t.Timeout + t.Timeout/3
						err := p.Push(ctx, p.queue, t)
						if err != nil {
							logs.S.Warn(err)
						}
						logs.S.Warnw("处理失败，进行重试", "queue", p.queue, "err", err)
					} else {
						logs.S.Errorw("处理失败，任务丢弃", "queue", p.queue, "err", err)
						if p.CallBack != nil {
							p.CallBack(t, err)
						}
					}
				} else {
					logs.S.Infow("处理成功", "queue", p.queue)
					if p.CallBack != nil {
						p.CallBack(t, err)
					}
				}

			}

		}()

	}

}

func (p *CrawlerEngine) Push(ctx context.Context, queue string, task task.Task) error {
	logs.S.Infow("下发任务", "global_id", task.GlobalId, "url", task.CrawlerUrl)
	taskStr, _ := json.Marshal(task)
	return p.redis.RPush(ctx, queue, taskStr).Err()
}
func (p *CrawlerEngine) NewMutex() *sync.Mutex {
	return sync.New(p.redis)
}
func (p *CrawlerEngine) Close() error {
	return p.redis.Close()
}
