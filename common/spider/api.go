package spider

import (
	"bytes"
	"crypto/tls"
	"go.uber.org/zap"
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
				Body: body,
				SpiderType: a.spiderType,
			}
			return resp, nil
		} else {
			zap.S().Warnf("retry request err", "url", request.Url, "err", err.Error(), "retry", i+1)
		}
	}
	return Response{}, err
}

func (a *ApiSpider) getResponse(request Request) ([]byte, error) {
	var netTransport *http.Transport
	if request.ProxyCallBack != nil {
		p := request.ProxyCallBack()
		proxy, err := url.Parse(p)
		if err != nil {
			return nil, err
		}
		netTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
			TLSClientConfig:&tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}
	client := &http.Client{Timeout: request.Timeout, Transport: netTransport}
	req, err := http.NewRequest(request.Param, request.Url, bytes.NewReader([]byte(request.Param)))
	if err != nil {
		zap.S().Errorf("请求数据出错", "error", err.Error())
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
		zap.S().Errorf("请求超时", "error", err.Error())
		return nil, err
	}
	defer response.Body.Close()
	data, err := ReadAll(response.Body, a.size)
	if err != nil {
		zap.S().Errorf("读取响应数据出错", "err:", err.Error())
		return nil, err
	}
	return data, nil
}