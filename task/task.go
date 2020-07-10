package task

import (
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
	Header        map[string]interface{} `json:"header"`
	Retry         int                    `json:"retry"`
	Total         int                    `json:"total"`
	Timeout       time.Duration          `json:"timeout"`
	Time          int64                  `json:"time"`
}
