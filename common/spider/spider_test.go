package spider

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestNewShtmlSpider(t *testing.T) {
	sp := NewCloudflareSpider(1 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:    "https://rarbgproxied.org/torrents.php?category=2;18;41;49&page=1",
		Method: "GET",
		Headers: map[string]string{
			"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			//"accept-encoding":           "gzip, deflate, br",
			"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
			"cache-control":             "no-cache",
			"cookie":                    "__cfduid=d74a2ce216735c9c3824f9d8b6630a2ce1605234630; gaDts48g=q8h5pp9t; tcc; aby=2; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; gaDts48g=q8h5pp9t; cf_clearance=b92bd96a8d21feeabe61ac4bf6fe0a70493944af-1605681737-0-150; ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; skt=IGX6Uufhj9; skt=IGX6Uufhj9; expla=2; ppu_sub_9ef78edf998c4df1e1636c9a474d9f47=2; ppu_delay_9ef78edf998c4df1e1636c9a474d9f47=1",
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
