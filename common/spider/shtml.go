package spider

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/SunMaybo/jewel-crawler/common/spider/charset"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	"gopkg.in/resty.v1"
	"net/http"
	"regexp"
	"strings"
)

type ShtmlSpider struct {
	spiderType SpiderType
	size       int
}

//请求数据最大size限制
func NewShtmlSpider(size int) *ShtmlSpider {
	return &ShtmlSpider{
		spiderType: Shtml,
		size:       size,
	}
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
			logs.S.Errorw("请求数据出错", "error", err.Error(), "retry", i+1)
			continue
		}
		if resp.StatusCode() >= 200 {
			readerCloser := resp.RawBody()
			defer readerCloser.Close()
			var buff []byte
			buff, err = ReadAll(readerCloser, s.size)
			if err != nil {
				logs.S.Errorw("读取响应数据出错", "err:", err.Error(), "retry", i+1)
				continue
			}
			encode := getResponseCharset(resp)
			return Response{
				RedirectUrl: resp.RawResponse.Request.URL.String(),
				charset:     encode,
				body:        charset.MustDecodeBytes(buff, encode),
				SpiderType:  s.spiderType,
			}, nil
		} else {
			return Response{}, errors.New(fmt.Sprintf("shtml rquest err statusCode:%d", resp.StatusCode()))
		}
	}
	return Response{}, err
}

func (s *ShtmlSpider) getResponse(request Request) (*resty.Response, error) {
	client := resty.New()
	if request.ProxyCallBack != nil {
		proxy := request.ProxyCallBack()
		if proxy != "" {
			client.SetProxy(request.ProxyCallBack())
		}
	}
	if request.SocketProxyCallBack != nil {
		// create a dialer
		u, p, a := request.SocketProxyCallBack()
		if a != "" {
			dialer, err := proxy.SOCKS5("tcp", a, &proxy.Auth{User: u, Password: p}, proxy.Direct)
			if err != nil {
				logs.S.Error("Unable to obtain proxy dialer: %v\n", err)
			}

			// create a transport
			ptransport := &http.Transport{Dial: dialer.Dial}
			// set transport into resty
			resty.SetTransport(ptransport)
		}

	}
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
	client.SetTimeout(request.Timeout)
	client.SetRetryCount(0)
	client.SetDoNotParseResponse(true)
	if request.CookieCallBack != nil {
		client.SetCookies(request.CookieCallBack())
	}
	r := client.R()
	if request.Headers != nil {
		isUserAgent := false
		for k, v := range request.Headers {
			if strings.Contains(v, "User-Agent") {
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
