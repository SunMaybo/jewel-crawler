package spider

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type FileSpider struct {
	disableKeepAlive bool
	spiderType       SpiderType
	size             int
	Jar              http.CookieJar
}

var fileSpiderLock sync.Mutex

var fileSpider *FileSpider

//请求数据最大size限制
func NewFileSpider(size int) *FileSpider {
	return &FileSpider{
		spiderType:       File,
		size:             size,
		disableKeepAlive: true,
	}
}
func NewSingleFileSpider(size int) *FileSpider {
	if fileSpider != nil {
		return fileSpider
	}
	fileSpiderLock.Lock()
	defer fileSpiderLock.Unlock()
	fileSpider = &FileSpider{
		spiderType:       File,
		size:             size,
		disableKeepAlive: false,
	}
	return fileSpider
}
func (f *FileSpider) Do(request Request) (Response, error) {
	if request.Retry <= 0 {
		request.Retry = 3
	}
	var err error
	for i := 0; i < request.Retry; i++ {
		var body []byte
		body, err = f.getResponse(request)
		if err == nil {
			resp := Response{
				body:       body,
				SpiderType: f.spiderType,
				charset:    recognitionFileFormat(body),
			}
			return resp, nil
		} else {
			logs.S.Warnw("redirect request err", "url", request.Url, "err", err.Error(), "redirect", i+1)
		}
	}
	return Response{}, err
}

func (f *FileSpider) getResponse(request Request) ([]byte, error) {

	netTransport := &http.Transport{
		DisableKeepAlives: f.disableKeepAlive,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}
	if request.ProxyCallBack != nil {
		p := request.ProxyCallBack()
		if p != "" {
			proxy, err := url.Parse(p)
			if err != nil {
				return nil, err
			}
			netTransport.Proxy = http.ProxyURL(proxy)
		}
	}
	if request.SocketProxyCallBack != nil {
		user, pwd, url := request.SocketProxyCallBack()

		if url != "" && user != "" && pwd != "" {
			auth := proxy.Auth{
				User:     user,
				Password: pwd,
			}
			// 设置代理
			dialer, err := proxy.SOCKS5("tcp", url, &auth, proxy.Direct)
			if err != nil {
				logs.S.Error(err)
				return nil, err
			}
			netTransport.Dial = dialer.Dial
		} else if url != "" {
			// 设置代理
			dialer, err := proxy.SOCKS5("tcp", url, nil, proxy.Direct)
			if err != nil {
				logs.S.Error(err)
				return nil, err
			}
			netTransport.Dial = dialer.Dial
		}
	}
	client := &http.Client{Timeout: request.Timeout, Transport: netTransport}
	if f.Jar != nil {
		client.Jar = f.Jar
	}
	req, err := http.NewRequest(request.Method, request.Url, bytes.NewReader([]byte(request.Param)))
	if err != nil {
		logs.S.Errorw("请求数据出错", "error", err.Error())
		return nil, err
	}
	if request.CookieCallBack != nil {
		for _, cookie := range request.CookieCallBack() {
			req.AddCookie(cookie)
		}

	}
	if request.Headers != nil {
		isUserAgent := false
		for k, v := range request.Headers {
			if strings.Contains(k, "User-Agent") {
				isUserAgent = true
			}
			req.Header.Add(k, v)
		}
		if !isUserAgent {
			req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0")
		}
	} else {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0")
	}
	req.Header.Add("accept-encoding","gzip, deflate, br")
	response, err := client.Do(req)
	if err != nil {
		logs.S.Errorw("请求超时", "error", err.Error())
		return nil, err
	}
	buff, err := UnZipHttpResp(response, f.size)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

var imageTypeMap = map[string]string{
	"0d0a1a0a": "png",
	"ff":       "jpg",
	"424d":     "bmp",
	"474946":   "gif",
	"574542":   "webp",
	"0000":     "ico",
	"42":       "tiff",
	"4D4D":     "tif",
	"4949":     "tif",
}

// 识别file格式
func recognitionFileFormat(content []byte) string {
	sliceContent1 := content[4:8]
	sliceContent2 := content[0:1]
	sliceContent3 := content[0:2]
	sliceContent4 := content[0:3]
	sliceContent5 := content[8:11]
	sliceContent6 := content[0:2]
	sliceContent7 := content[2:3]
	sliceContent8 := content[0:2]

	var content_list = [...][]byte{sliceContent1, sliceContent2, sliceContent3, sliceContent4, sliceContent5, sliceContent6, sliceContent7, sliceContent8}

	for _, v := range content_list {
		encodedStr := hex.EncodeToString(v)
		imageType, ok := imageTypeMap[encodedStr]
		if ok {
			return imageType
		}
	}
	logs.S.Infow(" set png type as default")
	return ""
}
