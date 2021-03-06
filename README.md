# jewl-crawler 简介

jewl-crawler 是一个分布式爬虫框架

jewel-crawler 可以用来网站数据采集，支持通过接口、静态页面、动态页面、以及图片等资源抓取任务，通过Redis进行任务的流转，支持多层次抓取需求，比如支持新闻类网站以列表和详情页两层数据流转采集。

架构集成了[GraphQuery](https://github.com/storyicon/graphquery)方便对数据进行解析，[go-readability](https://github.com/go-shiori/go-readability)适用于大部分的文章采集。



# 安装与使用

go get -u github.com/SunMaybo/jewel-crawler





# 栗子
## 项目启动
通过全局SetLogLevel方法设置日志级别

使用[go-redis](github.com/go-redis/redis/v8)启动Redis，Concurrent为支持最大并发数量，防止因为goroutine开启过大造成内存开销巨大。

Queue为使用的队列名字，实践中我们通过依据：图片、视频、文档、静态页面、动态页面、接口等对任务进行划分，依据是请求耗时和数据大小。
```
	jewel_crawler.SetLogLevel("info")
	engine := jewel_crawler.New(&jewel_crawler.Config{
		Queue:      config.Queue,
		Concurrent: config.Concurrent,
		Redis:      &redis.Options{
			Addr:         config.Redis.Addr,
			MinIdleConns: config.Redis.Idle,
			PoolSize:     config.Redis.Active,
			DB:           config.Redis.DB,
			Password:     config.Redis.Password,
			Username:     config.Redis.Name,
		},
	})
```

## 注册抓取模版
```
engine.Pipeline.AddCrawler("default_html_crawler", &crawler.DefaultHtmlCrawler{})
```

## 实现抓取模版
抓取模版需要实现接口方法并注册
```
type Crawler interface {
	Collect(event CollectEvent) (string, error)
	Parser(event ParserEvent) (map[string]interface{}, error)
	Storage(event StorageEvent) error
}
```
## 文章抓取的模版
其中event中封装了抓取组建，并且可以制定抓取数据最大大小，防止数据过大撑爆内存
```
type DefaultHtmlCrawler struct {
}

func (dhc *DefaultHtmlCrawler) Collect(event CollectEvent) (string, error) {
	s := spider.NewShtmlSpider(1 * 1024 * 1024)
	resp, err := s.Do(spider.Request{
		Url:     event.Task.CrawlerUrl,
		Method:  event.Task.Method,
		Param:   event.Task.Param,
		Headers: event.Task.Header,
		Timeout: event.Task.Timeout,
	})
	if err != nil {
		return "", err
	}
	return resp.GetContent(), nil
}
func (dhc *DefaultHtmlCrawler) Parser(event ParserEvent) (map[string]interface{}, error) {
	data, err := event.ReadabilityParser(event.Content, event.Task.CrawlerUrl)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (dhc *DefaultHtmlCrawler) Storage(event StorageEvent) error {
	logs.S.Info(event.Data)
	return nil
}
```
## 抓取组件

### 接口采集

```
apiSpider := event.ApiSpider(1024 * 1024)
	resp, err := apiSpider.Do(spider.Request{
		Url:     event.Task.CrawlerUrl,
		Method:  event.Task.Method,
		Param:   event.Task.Param,
		Headers: event.Task.Header,
		Timeout: event.Task.Timeout,
		ProxyCallBack: func() string {
			return "http://127.0.0.1:7890"
		},
	})
```
### 静态页面采集
静态页面采集返回的数据进行自动转码UTF-8，解决不同网站编码不同带来乱码问题，以及返回RedirectUrl最大重定向深度为10
```
shtmlSpider := event.ShtmlSpider(1024 * 1024)
	resp, err := shtmlSpider.Do(spider.Request{
		Url:    event.Task.CrawlerUrl,
		Method: "GET",
		ProxyCallBack: func() string {
			return proxy.GetRandomProxy()
		},
	})
    fmt.Println(resp.GetContent())
	fmt.Println(resp.GetCharset())
	fmt.Println(resp.RedirectUrl)
```

### 动态页面采集
动态采集使用[chromedp](github.com/chromedp/chromedp)底层依赖于Chrome浏览器需要在启动Chrome浏览器的环境中进行启用。
```
dhtmlSpider := event.DhtmlSpider(1024 * 1024)
	resp, err := dhtmlSpider.Do(spider.Request{
		Url:    event.Task.CrawlerUrl,
		Method: "GET",
		ProxyCallBack: func() string {
			return proxy.GetRandomProxy()
		},
	})
```

### 文件采集
目前仅支持图片采集自动解析图片格式
```
sp := spider.NewFileSpider(50 * 1024 * 1024)
	resp, err := sp.Do(spider.Request{
		Url:     event.Task.CrawlerUrl,
		Method:  "GET",
		Timeout: 60 * time.Second,
		ProxyCallBack: func() string {
			return proxy.GetRandomProxy()
		},
	})
	
fmt.Println(resp.GetCharset()) //图片格式
```
## 任务下发
其中TinyExtras 用户任务小数据的转移，较大数据临时存储通过temp_storage_id 进行关联，数据存储在redis中
```
childImage := task.ChildTask{
			CrawlerName: ImgCrawler,
			CrawlerUrl:  img,
			ContentType: "json",
			TempStorageId: "",
			Method:      "GET",
			Index:       i,
			TinyExtras: map[string]interface{}{
				"path": p.ImagePath(img),
			},
		}
event.Task.Next(context.Background(), event.Queue, childImage)
```
### 较大数据临时存储
```
event.TempStorage.Set(context.Background(),"id",temp.Temp{},1*time.Second)
event.TempStorage.Get(context.Background(),"id")
event.TempStorage.Clear(context.Background(),"id")
```

## 使用[GraphQuery](https://github.com/storyicon/graphquery)进行解析

```
patternTotal := "{\n  total `css(\".cta-link\");regex(\"[0-9]+\")`\n}"
totalSecondInter, _ := event.Parser(respTotal.GetContent(), patternTotal)
```

## 启动任务
```
config.CrawlerEngine.Start(context.Background(), config.MaxRetry)
```

## 其它
### 分布式锁支持
#### 阻塞锁
```
lock := s.Config.CrawlerEngine.NewMutex()
lock.Name = "lock"   \\锁名字
lock.Timeout = 1*time.Second  \\超时时间设置
lock.Lock()
defer lock.UnLock()
```

### 非阻塞锁

```
lock := s.Config.CrawlerEngine.NewMutex()
lock.Name = "lock"   \\锁名字
lock.Timeout = 1*time.Second  \\超时时间设置
lock.NLock()
defer lock.UnLock()
```
### 全局偏移量存储
```
lock := s.Config.CrawlerEngine.NewMutex()
lock.Name = "lock"   \\锁名字
lock.SetOffset("1")
offset,err:=lock.GetOffset()
```



