package spider

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type FileSpider struct {
	spiderType SpiderType
	size       int
}

//请求数据最大size限制
func NewFileSpider(size int) *FileSpider {
	return &FileSpider{
		spiderType: File,
		size:       size,
	}
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
			logs.S.Warnw("retry request err", "url", request.Url, "err", err.Error(), "retry", i+1)
		}
	}
	return Response{}, err
}

func (f *FileSpider) getResponse(request Request) ([]byte, error) {
	var netTransport *http.Transport
	netTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
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
		auth := proxy.Auth{
			User:     user,
			Password: pwd,
		}
		if url != "" {
			// 设置代理
			dialer, err := proxy.SOCKS5("tcp", url, &auth, proxy.Direct)
			if err != nil {
				logs.S.Error(err)
				return nil, err
			}
			netTransport.Dial = dialer.Dial
		}
	}
	client := &http.Client{Timeout: request.Timeout, Transport: netTransport}
	if request.CookieJarCallBack != nil {
		jar, _ := cookiejar.New(nil)
		u, _ := url.Parse(request.Url)
		jar.SetCookies(u, request.CookieJarCallBack())
		client.Jar = jar
	}
	req, err := http.NewRequest(request.Method, request.Url, bytes.NewReader([]byte(request.Param)))
	if err != nil {
		logs.S.Errorw("请求数据出错", "error", err.Error())
		return nil, err
	}

	if request.Headers != nil {
		isUserAgent := false
		for k, v := range request.Headers {
			if strings.Contains(v, "User-Agent") {
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
	response, err := client.Do(req)
	if err != nil {
		logs.S.Errorw("请求超时", "error", err.Error())
		return nil, err
	}
	defer response.Body.Close()
	data, err := ReadAll(response.Body, f.size)
	if err != nil {
		logs.S.Errorw("读取响应数据出错", "err:", err.Error())
		return nil, err
	}
	return data, nil
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
