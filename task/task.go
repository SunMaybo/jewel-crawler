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
	Redis         *redis.Client          `json:"-"`
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
	Header        map[string]string      `json:"header"`
}

func (t *Task) Next(ctx context.Context, queue string, child ChildTask) error {
	task := Task{
		ParentId:      t.TaskId,
		CrawlerName:   child.CrawlerName,
		TaskId:        common.GenerateRandomID(),
		GlobalId:      t.GlobalId,
		Time:          time.Now().Unix(),
		CrawlerUrl:    child.CrawlerUrl,
		TempStorageId: child.TempStorageId,
		Depth:         t.Depth + 1,
		ContentType:   child.ContentType,
		Header:        child.Header,
		Index:         child.Index,
		Method:        child.Method,
		Param:         child.Param,
		TinyExtras:    child.TinyExtras,
		Website:       t.Website,
		Timeout:       t.Timeout,
	}
	buff, _ := json.Marshal(task)
	return t.Redis.LPush(ctx, queue, string(buff)).Err()
}
