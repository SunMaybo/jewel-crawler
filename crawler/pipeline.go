package crawler

import (
	"context"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type PipeLine struct {
	filterFunc   map[FilterType]func(filter Filter) bool
	reportFunc   map[ReportType]func(report Report) error
	crawlerInter map[task.CrawlerName]Crawler
	tempStorage  *temp.TempStorage
}

func New(tempStorage *temp.TempStorage) *PipeLine {
	return &PipeLine{
		filterFunc:   make(map[FilterType]func(filter Filter) bool),
		reportFunc:   make(map[ReportType]func(report Report) error),
		crawlerInter: make(map[task.CrawlerName]Crawler),
		tempStorage:  tempStorage,
	}
}

//选择抓取模版、数据抓取、report、过滤

type FilterType int

type ReportType int

const (
	CrawlerBeforeFilter FilterType = iota + 1
	CrawlerAfterFilter
	ParserBeforeFilter
	ParserAfterFilter
	StorageBeforeFilter
	StorageAfterFilter

	CrawlerBeforeReport ReportType = iota + 1
	CrawlerAfterReport
	ParserBeforeReport
	ParserAfterReport
	StorageBeforeReport
	StorageAfterReport
)

func (f FilterType) String() string {
	switch f {
	case CrawlerBeforeFilter:
		return "crawler-before-filter"
	case CrawlerAfterFilter:
		return "crawler-after-filter"
	case ParserBeforeFilter:
		return "parser-before-filter"
	case ParserAfterFilter:
		return "parser-after-filter"
	case StorageBeforeFilter:
		return "storage-before-filter"
	case StorageAfterFilter:
		return "storage-after-filter"
	default:
		return ""
	}

}

func (r ReportType) String() string {
	switch r {
	case CrawlerBeforeReport:
		return "crawler-before-report"
	case CrawlerAfterReport:
		return "crawler-after-report"
	case ParserBeforeReport:
		return "parser-before-report"
	case ParserAfterReport:
		return "parser-after-report"
	case StorageBeforeReport:
		return "storage-before-report"
	case StorageAfterReport:
		return "storage-after-report"
	default:
		return ""
	}

}

type Filter struct {
	Task       task.Task
	Temp       temp.Temp
	Data       map[string]interface{}
	FilterType FilterType
}
type Report struct {
	Task       task.Task
	Error      error
	ReportType ReportType
}

//添加过滤器
func (p *PipeLine) AddFilter(filterType FilterType, filterFunc func(filter Filter) bool) {
	p.filterFunc[filterType] = filterFunc
}

//添加上报
func (p *PipeLine) AddReport(reportType ReportType, reportFunc func(report Report) error) {
	p.reportFunc[reportType] = reportFunc
}

//添加抓取模版
func (p *PipeLine) AddCrawler(crawlerName task.CrawlerName, crawler Crawler) {
	p.crawlerInter[crawlerName] = crawler
}

func (p *PipeLine) Invoke(ctx context.Context, task task.Task) error {
	start := time.Now()
	if crawler, ok := p.crawlerInter[task.CrawlerName]; ok {
		var temp temp.Temp
		var err error
		//获取临时存储数据
		tsid := task.TempStorageId
		if tsid != "" {
			temp, err = p.tempStorage.Get(ctx, tsid)
			if err != nil && redis.Nil != err {
				zap.S().Errorf("redis err", "message", err.Error())
				return err
			}
		}

		//抓取前置过滤
		if filterFunc, ok := p.filterFunc[CrawlerBeforeFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				FilterType: CrawlerBeforeFilter,
			}) {
				return nil
			}

		}
		//前置上报
		if reportFunc, ok := p.reportFunc[CrawlerBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: CrawlerBeforeReport,
			})
			if err != nil {
				zap.S().Warnf("crawler before report err", "message", err.Error())
			}

		}
		content, err := crawler.Collect(task, temp)

		//后置上报
		if reportFunc, ok := p.reportFunc[CrawlerAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				Error:      err,
				ReportType: CrawlerAfterReport,
			})
			if err != nil {
				zap.S().Warnf("crawler after report err", "message", err.Error())
			}

		}

		if err != nil {
			zap.S().Warnf("crawler err", "crawler_name", task.CrawlerName, "global_id",
				task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId,
				"crawler_url", task.CrawlerUrl, "interval", time.Now().Sub(start).Milliseconds(), "message", err.Error())
			return err
		}

		zap.S().Infof("crawler success", "crawler_name", task.CrawlerName, "global_id", task.GlobalId,
			"parent_id", task.ParentId, "task_id", task.TaskId, "crawler_url",
			task.CrawlerUrl, "interval", time.Now().Sub(start).Milliseconds())

		//抓取后置过滤
		if filterFunc, ok := p.filterFunc[CrawlerAfterFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				FilterType: CrawlerAfterFilter,
			}) {
				return nil
			}

		}

		//解析前置过滤
		if filterFunc, ok := p.filterFunc[ParserBeforeFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				FilterType: ParserBeforeFilter,
			}) {
				return nil
			}

		}
		//解析前置上报
		if reportFunc, ok := p.reportFunc[ParserBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: ParserBeforeReport,
			})
			if err != nil {
				zap.S().Warnf("parser before report err", "message", err.Error())
			}

		}

		//数据解析
		data, err := crawler.Parser(task, temp, content)

		//解析后置上报
		if reportFunc, ok := p.reportFunc[ParserAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				Error:      err,
				ReportType: ParserAfterReport,
			})
			if err != nil {
				zap.S().Warnf("parser after report err", "message", err.Error())
			}

		}

		if err != nil {
			zap.S().Errorf("parser err", "crawler_name", task.CrawlerName, "global_id", task.GlobalId,
				"parent_id", task.ParentId, "task_id", task.TaskId, "crawler_url", task.CrawlerUrl,
				"interval", time.Now().Sub(start).Milliseconds(), "message", err.Error())
			return err
		}

		zap.S().Infof("parser success", "crawler_name", task.CrawlerName, "global_id", task.GlobalId,
			"parent_id", task.ParentId, "task_id", task.TaskId, "crawler_url", task.CrawlerUrl,
			"interval", time.Now().Sub(start).Milliseconds())

		//解析后置过滤
		if filterFunc, ok := p.filterFunc[ParserAfterFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				Data:       data,
				FilterType: ParserAfterFilter,
			}) {
				return nil
			}

		}

		//存储前置过滤
		if filterFunc, ok := p.filterFunc[StorageBeforeFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				Data:       data,
				FilterType: StorageBeforeFilter,
			}) {
				return nil
			}

		}
		//存储前置上报
		if reportFunc, ok := p.reportFunc[StorageBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: StorageBeforeReport,
			})
			if err != nil {
				zap.S().Warnf("storage before report err", "message", err.Error())
			}

		}

		err = crawler.Storage(task, data)

		//存储后置上报
		if reportFunc, ok := p.reportFunc[StorageAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: StorageAfterReport,
				Error:      err,
			})
			if err != nil {
				zap.S().Warnf("storage after report err", "message", err.Error())
			}

		}
		if err != nil {
			zap.S().Warnf("storage err", "crawler_name", task.CrawlerName, "global_id",
				task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId,
				"crawler_url", task.CrawlerUrl, "interval", time.Now().Sub(start).Milliseconds(), "message", err.Error())
			return err
		}
		zap.S().Infof("storage success", "crawler_name", task.CrawlerName, "global_id",
			task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId,
			"crawler_url", task.CrawlerUrl, "interval", time.Now().Sub(start).Milliseconds())
		//存储后置过滤
		if filterFunc, ok := p.filterFunc[StorageAfterFilter]; ok {
			if filterFunc(Filter{
				Task:       task,
				Temp:       temp,
				Data:       data,
				FilterType: StorageAfterFilter,
			}) {
				return nil
			}
		}
		return nil
	} else {
		zap.S().Warnf("no grab template or grab name is illegal", "crawler_name", task.CrawlerName)
		return nil
	}
	return nil
}
