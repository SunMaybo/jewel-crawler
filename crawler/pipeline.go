package crawler

import (
	"context"
	"github.com/SunMaybo/jewel-crawler/logs"
	"github.com/SunMaybo/jewel-crawler/task"
	"github.com/SunMaybo/jewel-crawler/temp"
	"time"
)

type PipeLine struct {
	reportFunc   map[ReportType]func(report Report) error
	crawlerInter map[task.CrawlerName]Crawler
	tempStorage  *temp.TempStorage
	queue        string
}

func New(queue string, tempStorage *temp.TempStorage) *PipeLine {
	return &PipeLine{
		reportFunc:   make(map[ReportType]func(report Report) error),
		crawlerInter: make(map[task.CrawlerName]Crawler),
		tempStorage:  tempStorage,
		queue:        queue,
	}
}

//选择抓取模版、数据抓取、report、过滤

type ReportType int

const (
	CrawlerBeforeReport ReportType = iota + 1
	CrawlerAfterReport
	ParserBeforeReport
	ParserAfterReport
	StorageBeforeReport
	StorageAfterReport
)

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

type Report struct {
	Task       task.Task
	Error      error
	ReportType ReportType
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
		var err error
		//前置上报
		if reportFunc, ok := p.reportFunc[CrawlerBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: CrawlerBeforeReport,
			})
			if err != nil {
				logs.S.Warnw("crawler before report err", "message", err.Error())
			}

		}
		content, err := crawler.Collect(CollectEvent{
			Task:  task,
			Queue: p.queue,
			Event: Event{
				TempStorage: p.tempStorage,
			},
		})

		//后置上报
		if reportFunc, ok := p.reportFunc[CrawlerAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				Error:      err,
				ReportType: CrawlerAfterReport,
			})
			if err != nil {
				logs.S.Warnw("crawler after report err", "message", err.Error())
			}

		}

		if err != nil {
			logs.S.Errorw("crawler err", "crawler_name",task.CrawlerName, "err", err,  "crawler_url", task.CrawlerUrl, "global_id",
				task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId, "interval", time.Since(start).String())
			return err
		}

		logs.S.Infow("crawler success", "crawler_name", task.CrawlerName, "global_id", task.GlobalId,
			"parent_id", task.ParentId, "task_id", task.TaskId, "crawler_url",
			task.CrawlerUrl, "interval", time.Since(start).String())

		//解析前置上报
		if reportFunc, ok := p.reportFunc[ParserBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: ParserBeforeReport,
			})
			if err != nil {
				logs.S.Warnw("parser before report err", "message", err.Error())
			}

		}

		//数据解析
		data, err := crawler.Parser(ParserEvent{
			Task:    task,
			Queue:   p.queue,
			Content: content,
			Event: Event{
				TempStorage: p.tempStorage,
			},
		})

		//解析后置上报
		if reportFunc, ok := p.reportFunc[ParserAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				Error:      err,
				ReportType: ParserAfterReport,
			})
			if err != nil {
				logs.S.Warnw("parser after report err", "message", err.Error())
			}

		}

		if err != nil {
			logs.S.Errorw("parser err", "crawler_name", task.CrawlerName, "err", err, "crawler_url", task.CrawlerUrl, "global_id", task.GlobalId,
				"parent_id", task.ParentId, "task_id", task.TaskId,
				"interval", time.Since(start).String())
			return err
		}

		if data != nil && len(data) > 0 {
			logs.S.Infow("parser success", "crawler_name", task.CrawlerName, "global_id", task.GlobalId,
				"parent_id", task.ParentId, "task_id", task.TaskId, "crawler_url", task.CrawlerUrl,
				"interval", time.Since(start).String())
		}

		//存储前置上报
		if reportFunc, ok := p.reportFunc[StorageBeforeReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: StorageBeforeReport,
			})
			if err != nil {
				logs.S.Warnw("storage before report err", "message", err.Error())
			}

		}

		err = crawler.Storage(StorageEvent{
			Task:  task,
			Data:  data,
			Queue: p.queue,
			Event: Event{
				TempStorage: p.tempStorage,
			},
		})

		//存储后置上报
		if reportFunc, ok := p.reportFunc[StorageAfterReport]; ok {
			err := reportFunc(Report{
				Task:       task,
				ReportType: StorageAfterReport,
				Error:      err,
			})
			if err != nil {
				logs.S.Warnw("storage after report err", "message", err.Error())
			}

		}
		if err != nil {
			logs.S.Errorw("storage err", "crawler_name", task.CrawlerName,"err", err, "crawler_url", task.CrawlerUrl, "global_id",
				task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId,
				"interval", time.Since(start).String())
			return err
		}
		logs.S.Infow("storage success", "crawler_name", task.CrawlerName, "global_id",
			task.GlobalId, "parent_id", task.ParentId, "task_id", task.TaskId,
			"crawler_url", task.CrawlerUrl, "interval", time.Since(start).String())
		return nil
	} else {
		logs.S.Warnw("no grab template or grab name is illegal", "crawler_name", task.CrawlerName)
		return nil
	}
	return nil
}
