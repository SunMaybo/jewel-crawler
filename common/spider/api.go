package spider

import (
	"bytes"
	"crypto/tls"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var apiLock sync.Mutex
var apiSpider *ApiSpider

type ApiSpider struct {
	spiderType       SpiderType
	size             int
	disableKeepAlive bool
	client           *http.Client
	Jar              http.CookieJar
}

//请求数据最大size限制
func NewApiSpider(size int) *ApiSpider {
	return &ApiSpider{
		spiderType:       Api,
		size:             size,
		disableKeepAlive: true,
	}
}

func NewSingleApiSpider(size int, timeout time.Duration, transport *http.Transport) *ApiSpider {
	if apiSpider != nil {
		return apiSpider
	}
	apiLock.Lock()
	defer apiLock.Unlock()
	apiSpider = &ApiSpider{
		spiderType:       Api,
		size:             size,
		disableKeepAlive: false,
	}
	if transport != nil {
		transport = &http.Transport{}
	}
	client := &http.Client{Timeout: timeout, Transport: transport}
	apiSpider.client = client
	return apiSpider
}
func (a *ApiSpider) Do(request Request) (Response, error) {
	if request.Retry <= 0 {
		request.Retry = 3
	}
	var err error
	for i := 0; i < request.Retry; i++ {
		var body []byte
		body, err = a.getResponse(request)
		if err == nil {
			resp := Response{
				body:       body,
				SpiderType: a.spiderType,
			}
			return resp, nil
		} else {
			logs.S.Warnw("retry request err", "url", request.Url, "err", err.Error(), "retry", i+1)
		}
	}
	return Response{}, err
}

func (a *ApiSpider) getClient(request Request) (*http.Client, error) {
	if !a.disableKeepAlive {
		return a.client, nil
	}
	netTransport := &http.Transport{
		DisableKeepAlives: true,
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
			// 设置代理
			auth := proxy.Auth{
				User:     user,
				Password: pwd,
			}
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
	if a.Jar != nil {
		client.Jar = a.Jar
	}
	return client, nil
}

func (a *ApiSpider) getResponse(request Request) ([]byte, error) {

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
	client, err := a.getClient(request)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(req)
	if err != nil {
		logs.S.Errorw("请求超时", "error", err.Error())
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.S.Errorw("读取响应数据出错", "err:", err.Error())
		return nil, err
	}
	return data, nil
}
