package spider

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestNewShtmlSpider(t *testing.T) {
	sp := NewShtmlSpider(5 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:    "https://m.kooora.com/?n=990171&o=n",
		Method: "GET",
		Headers: map[string]string{
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"accept-encoding":           "gzip, deflate, br",
			"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
			"cache-control":             "no-cache",
			"pragma":                    "no-cache",
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
		},
		ProxyCallBack: func() string {
			return "http://127.0.0.1:7890"
		},
		Timeout: 30 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetContent())
	t.Log(resp.RedirectUrl)

}

func TestNewApiSpider(t *testing.T) {
	sp := NewCloudflareSpider(5 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:    "https://www1.yesmovies.so/movie/filter/movie/all/all/all/all/most/?page=1",
		Method: "GET",
		Headers: map[string]string{
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
			"cache-control":             "no-cache",
			"pragma":                    "o-cache",
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "none",
			"sec-fetch-user":            "?1",
			"referer":                   "https://www1.yesmovies.so/movie/filter/movie/all/all/all/all/most/",
			"upgrade-insecure-requests": "1",
			"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36",
		},
		Timeout: 5 * time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetContent())

}

func TestNewFileSpider(t *testing.T) {
	sp := NewFileSpider(1 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:     "https://static001.infoq.cn/resource/image/3a/b3/3a1fa03a1b44ba5ee608680cbd3d28b3.png",
		Method:  "GET",
		Timeout: 5 * time.Second,
		ProxyCallBack: func() string {
			return "http://127.0.0.1:7890"
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetCharset())
	ioutil.WriteFile("a.jpg", resp.GetBytes(), 0775)

}
