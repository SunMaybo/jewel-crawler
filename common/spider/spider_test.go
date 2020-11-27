package spider

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestNewShtmlSpider(t *testing.T) {
	sp := NewCloudflareSpider(1 * 1024 * 1024)
	resp, err := sp.Do(Request{
		Url:    "https://www1.yesmovies.so/movie/filter/movie/all/all/all/all/most/",
		Method: "GET",
		Headers: map[string]string{
			"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			//"accept-encoding":           "gzip, deflate, br",
			"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
			"cache-control":   "no-cache",
			//	"cookie":                    "__cfduid=d74a2ce216735c9c3824f9d8b6630a2ce1605234630; gaDts48g=q8h5pp9t; tcc; aby=2; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; gaDts48g=q8h5pp9t; cf_clearance=b92bd96a8d21feeabe61ac4bf6fe0a70493944af-1605681737-0-150; ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; skt=IGX6Uufhj9; skt=IGX6Uufhj9; expla=2; ppu_sub_9ef78edf998c4df1e1636c9a474d9f47=2; ppu_delay_9ef78edf998c4df1e1636c9a474d9f47=1",
			//	"cookie":                    "ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; __cfduid=df720fa4237e9c930b7e19408685a046e1605838169; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; skt=USPF6qmyox; gaDts48g=q8h5pp9t; skt=USPF6qmyox; gaDts48g=q8h5pp9t; tcc; aby=2; expla=2",
			"cookie":                    "__cfduid=dc15260751b1abe389f349f5249776f0e1603261841; _ga=GA1.2.1555090272.1603261845; cf_clearance=35f0e0ce6f877abe217445716dea18a62680b94b-1605840245-0-150; gogoanime=6gd3see9sufj4t4j5ihhuqhfp2; _gid=GA1.2.730488476.1605840251; SL_GWPT_Show_Hide_tmp=1; SL_wptGlobTipTmp=1; _gat=1; AdskeeperStorage=%7B%220%22%3A%7B%7D%2C%22C900524%22%3A%7B%22page%22%3A3%2C%22time%22%3A1605840855294%7D%7D",
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
