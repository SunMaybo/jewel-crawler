package spider

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/SunMaybo/jewel-crawler/common/spider/charset"
	"github.com/SunMaybo/jewel-crawler/common/spider/redirect"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	resty "gopkg.in/resty.v1"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ShtmlSpider struct {
	spiderType       SpiderType
	size             int
	Jar              http.CookieJar
	disableKeepAlive bool
	client           *resty.Client
}

var shtmlSpiderLock sync.Mutex
var shtmlSpider *ShtmlSpider

//请求数据最大size限制
func NewShtmlSpider(size int) *ShtmlSpider {
	return &ShtmlSpider{
		spiderType:       Shtml,
		size:             size,
		disableKeepAlive: true,
	}
}
func NewSingleShtmlSpider(size int, timeout time.Duration, transport *http.Transport) *ShtmlSpider {
	if shtmlSpider != nil {
		return shtmlSpider
	}
	shtmlSpiderLock.Lock()
	defer shtmlSpiderLock.Unlock()
	shtmlSpider = &ShtmlSpider{
		spiderType:       Shtml,
		size:             size,
		disableKeepAlive: false,
	}
	if transport == nil {
		transport = &http.Transport{
		}
	}
	client := resty.NewWithClient(&http.Client{
		Transport: transport,
	})
	client.SetRedirectPolicy(redirect.FlexibleRedirectPolicy(20))
	client.SetTimeout(timeout)
	client.SetRetryCount(0)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetRetryMaxWaitTime(timeout / 3)
	shtmlSpider.client = client
	return shtmlSpider
}
func (s *ShtmlSpider) Do(request Request) (Response, error) {
	if request.Retry <= 0 {
		request.Retry = 3
	}
	var err error
	for i := 0; i < request.Retry; i++ {
		var resp *resty.Response
		resp, err = s.getResponse(request)
		if err != nil {
			logs.S.Errorw("请求数据出错", "error", err.Error(), "redirect", i+1)
			continue
		}
		if resp.StatusCode() >= 200 {
			redirectUrl := resp.RawResponse.Request.URL.String()
			encode := getResponseCharset(resp)
			var body []byte
			if filterNotHtml(encode) {
				body = []byte("content_type is illegal:" + encode)
			} else {
				readerCloser := resp.RawBody()
				buff, err := ReadAll(readerCloser, s.size)
				if err != nil {
					logs.S.Errorw("读取响应数据出错", "err:", err.Error(), "retry", i+1, "url", request.Url)
					continue
				}
				readerCloser.Close()
				body = charset.MustDecodeBytes(buff, encode)
			}
			return Response{
				RedirectUrl: redirectUrl,
				charset:     encode,
				body:        body,
				SpiderType:  s.spiderType,
				Cookies:     resp.Cookies(),
				Header:      resp.Header(),
			}, nil
		} else {
			return Response{}, errors.New(fmt.Sprintf("shtml rquest err statusCode:%d", resp.StatusCode()))
		}
	}
	return Response{}, err
}

func (s *ShtmlSpider) getClient(request Request) *resty.Client {
	if !s.disableKeepAlive {
		return s.client
	}
	client := resty.NewWithClient(&http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: s.disableKeepAlive,
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		},
	})
	client.SetDoNotParseResponse(true)
	if request.ProxyCallBack != nil {
		proxy := request.ProxyCallBack()
		if proxy != "" {
			client.SetProxy(request.ProxyCallBack())
		}
	}
	if request.SocketProxyCallBack != nil {
		// create a dialer
		u, p, a := request.SocketProxyCallBack()
		if a != "" && u != "" && p != "" {
			dialer, err := proxy.SOCKS5("tcp", a, &proxy.Auth{User: u, Password: p}, proxy.Direct)
			if err != nil {
				logs.S.Error("Unable to obtain proxy dialer: %v\n", err)
			}
			// create a transport
			ptransport := &http.Transport{Dial: dialer.Dial}
			// set transport into resty
			resty.SetTransport(ptransport)
		} else if a != "" {
			dialer, err := proxy.SOCKS5("tcp", a, nil, proxy.Direct)
			if err != nil {
				logs.S.Error("Unable to obtain proxy dialer: %v\n", err)
			}

			// create a transport
			ptransport := &http.Transport{Dial: dialer.Dial}
			// set transport into resty
			resty.SetTransport(ptransport)
		}

	}
	client.SetRedirectPolicy(redirect.FlexibleRedirectPolicy(20))
	client.SetTimeout(request.Timeout)
	client.SetRetryCount(0)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetRetryMaxWaitTime(request.Timeout / 3)
	if s.Jar != nil {
		client.SetCookieJar(s.Jar)
	}
	if request.CookieCallBack != nil {
		client.SetCookies(request.CookieCallBack())
	}
	return client
}

func (s *ShtmlSpider) getResponse(request Request) (*resty.Response, error) {
	client := s.getClient(request)
	r := client.R()
	if request.Headers != nil {
		isUserAgent := false
		for k, v := range request.Headers {
			if strings.Contains(k, "User-Agent") {
				isUserAgent = true
			}
			r.Header.Add(k, v)
		}
		if !isUserAgent {
			r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0")
		}
	} else {
		r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0")
	}
	resp, err := r.Get(request.Url)
	return resp, err
}

// 识别响应头编码方式
func getResponseCharset(response *resty.Response) string {
	charsetStrs := response.Header()["Content-Type"]
	if charsetStrs == nil {
		return ""
	}
	for _, charsetStr := range charsetStrs {
		if charsetStr != "" && strings.Index(charsetStr, "charset") != -1 {
			regexpArr := [2]string{"charset=(.*?);", "charset=(.*)"}
			for _, value := range regexpArr {
				reg := regexp.MustCompile(value) // 进行正则编译
				findResult := reg.FindStringSubmatch(charsetStr)
				if findResult != nil && len(findResult) > 1 && findResult[1] != "" && strings.Index("gbk,gb18030,gb2312,utf8,utf-8,ansi,big5,unicode,ascii", strings.ToLower(findResult[1])) != -1 {
					charsetStr = findResult[1]
					logs.S.Infow("response header find charset ------>", "charset", charsetStr)
					return charsetStr
				}
			}
		}
	}
	return ""
}
func filterNotHtml(contentType string) bool {
	contentTypes := []string{
		"text/html",
		"text/plain",
		"text/xml",
		"application/xhtml+xml",
		"application/xml",
		"application/atom+xml",
		"application/json",
		"utf-8",
	}
	if contentType == "" {
		return false
	}
	for _, s := range contentTypes {
		if strings.Contains(strings.ToLower(contentType), strings.ToLower(s)) {
			return false
		}
	}
	return true
}
