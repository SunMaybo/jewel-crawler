package spider

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/SunMaybo/jewel-crawler/logs"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
)

type ApiSpider struct {
	spiderType SpiderType
	size       int
}

//请求数据最大size限制
func NewApiSpider(size int) *ApiSpider {
	return &ApiSpider{
		spiderType: Api,
		size:       size,
	}
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

func (a *ApiSpider) getResponse(request Request) ([]byte, error) {
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
	if request.SocketProxyCallBack!=nil {
		user,pwd,url:=request.SocketProxyCallBack()

		auth := proxy.Auth{
			User:     user,
			Password: pwd,
		}

		// 设置代理
		dialer, err := proxy.SOCKS5("tcp", url, &auth, proxy.Direct)
		if err != nil {
			return nil,err
		}
		netTransport.DialContext= func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network,addr)
		}
	}
	client := &http.Client{Timeout: request.Timeout, Transport: netTransport}
	req, err := http.NewRequest(request.Method, request.Url, bytes.NewReader([]byte(request.Param)))
	if err != nil {
		logs.S.Errorw("请求数据出错", "error", err.Error())
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:76.0) Gecko/20100101 Firefox/76.0")
	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Add(k, v)
		}
	}
	response, err := client.Do(req)
	if err != nil {
		logs.S.Errorw("请求超时", "error", err.Error())
		return nil, err
	}
	defer response.Body.Close()
	data, err := ReadAll(response.Body, a.size)
	if err != nil {
		logs.S.Errorw("读取响应数据出错", "err:", err.Error())
		return nil, err
	}
	return data, nil
}
