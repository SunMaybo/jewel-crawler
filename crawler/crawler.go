package crawler

import (
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
)

type CollectEvent struct {
	Task task.Task
	Temp temp.Temp
}
type ParserEvent struct {
	Task    task.Task
	Temp    temp.Temp
	Content string
}
type StorageEvent struct {
	Task        task.Task
	Queue     string
	Data        map[string]interface{}
	TempStorage *temp.TempStorage
}

type Crawler interface {
	Collect(event CollectEvent) (string, error)
	Parser(event ParserEvent) (map[string]interface{}, error)
	Storage(event StorageEvent) error
}
