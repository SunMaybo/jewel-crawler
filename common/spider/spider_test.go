package spider

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestNewShtmlSpider(t *testing.T) {
	sp := NewShtmlSpider(1 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:     "https://java.ctolib.com/chromedp.html",
		Method:  "GET",
		Timeout: 5 * time.Second,
		ProxyCallBack: func() string {
			return "http://127.0.0.1:7890"
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetContent())
	t.Log(resp.GetCharset())

}

func TestNewApiSpider(t *testing.T) {
	sp := NewApiSpider(1 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:     "https://zxp.artron.net/specials/goods/index/?specialid=6231&auctionstatus=3&page=2&perpage=20&ordertype=1",
		Method:  "GET",
		Timeout: 5 * time.Second,
		ProxyCallBack: func() string {
			return "http://127.0.0.1:7890"
		},
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
