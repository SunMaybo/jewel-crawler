package crawler

import (
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
)

type Crawler interface {
	Collect(task task.Task, temp temp.Temp) (string, error)
	Parser(task task.Task, temp temp.Temp, content string) (map[string]interface{}, error)
	Storage(task task.Task, data map[string]interface{}) error
}
