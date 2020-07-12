package task

import (
	"context"
	"encoding/json"
	"github.com/SunMaybo/jewel-crawler/common"
	"github.com/go-redis/redis/v8"
	"time"
)

type CrawlerName string

type Task struct {
	CrawlerName   CrawlerName            `json:"crawler_name"`
	Website       string                 `json:"website"`
	GlobalId      string                 `json:"global_id"`
	ParentId      string                 `json:"parent_id"`
	TaskId        string                 `json:"task_id"`
	Depth         int                    `json:"depth"`
	Index         int                    `json:"index"`
	TempStorageId string                 `json:"temp_storage_id"`
	TinyExtras    map[string]interface{} `json:"tiny_extras"`
	CrawlerUrl    string                 `json:"crawler_url"`
	ContentType   string                 `json:"content_type"`
	Method        string                 `json:"method"`
	Param         string                 `json:"param"`
	Header        map[string]string      `json:"header"`
	Retry         int                    `json:"retry"`
	Total         int                    `json:"total"`
	Timeout       time.Duration          `json:"timeout"`
	Time          int64                  `json:"time"`
	redis         *redis.Client          `json:"-"`
}

type ChildTask struct {
	CrawlerName   CrawlerName            `json:"crawler_name"`
	Index         int                    `json:"index"`
	TempStorageId string                 `json:"temp_storage_id"`
	TinyExtras    map[string]interface{} `json:"tiny_extras"`
	CrawlerUrl    string                 `json:"crawler_url"`
	ContentType   string                 `json:"content_type"`
	Method        string                 `json:"method"`
	Param         string                 `json:"param"`
	Header        map[string]string `json:"header"`
}

func (t *Task) Next(ctx context.Context, queue string, child ChildTask) error {
	t.ParentId = t.TaskId
	t.TaskId = common.GenerateRandomID()
	t.Time = time.Now().Unix()
	t.Retry = 0
	t.CrawlerUrl = child.CrawlerUrl
	t.TempStorageId = child.TempStorageId
	t.CrawlerName = child.CrawlerName
	t.Depth += 1
	t.ContentType = child.ContentType
	t.Header = child.Header
	t.Index = child.Index
	t.Method = child.Method
	t.Param = child.Param
	t.TinyExtras = child.TinyExtras
	buff, _ := json.Marshal(t)
	return t.redis.LPush(ctx, queue, string(buff)).Err()
}
