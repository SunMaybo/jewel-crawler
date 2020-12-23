package spider

import (
	"bytes"
	"errors"
	"github.com/SunMaybo/jewel-crawler/common/parser"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type CloudflareSpider struct {
	disableKeepAlive bool
	spiderType       SpiderType
	size             int
	fileSpider       *FileSpider
	Jar              http.CookieJar
}

var cloudflareSpider *CloudflareSpider

//请求数据最大size限制
func NewCloudflareSpider(size int) *CloudflareSpider {

	return &CloudflareSpider{
		spiderType:       File,
		size:             size,
		disableKeepAlive: true,
		fileSpider:       NewFileSpider(size),
	}
}

func (f *CloudflareSpider) Do(request Request) (Response, error) {
	pattern := "{\n  url `css(\"form\");attr(\"action\")`\n  method `css(\"form\");attr(\"method\")`\n  params `css(\"input\")`[{\n   name `attr(\"name\")`\n   value `attr(\"value\")`\n\n}]\n\n}"
	shtml := NewShtmlSpider(f.size)
	resp, err := shtml.Do(request)
	if err != nil {
		return Response{}, err
	}
	if !strings.Contains(resp.GetContent(), "Your browser will redirect to your requested content shortly") {
		return resp, nil
	}
	inter, err := parser.Parser(string(resp.GetContent()), pattern)
	if err != nil {
		return Response{}, err
	}
	if inter == nil {
		return Response{}, errors.New("parser 5 second cloudflare err")
	}
	dataMap := inter.(map[string]interface{})
	u, err := url.Parse(request.Url)
	if err != nil {
		return Response{}, err
	}
	if dataMap["url"] == nil || dataMap["url"].(string) == "" {
		return Response{}, errors.New("parser 5 second cloudflare err")
	}
	url := dataMap["url"].(string)
	url = u.Scheme + "://" + u.Host + url
	var params []string
	if dataMap["params"] != nil {
		paramInters := dataMap["params"].([]interface{})
		for _, paramInter := range paramInters {
			paramMap := paramInter.(map[string]interface{})
			params = append(params, paramMap["name"].(string)+"="+paramMap["value"].(string))
		}
	}
	request.Param = strings.Join(params, "&")
	request.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	request.Url = url
	request.Method = dataMap["method"].(string)
	var cookies []string
	for _, cookie := range resp.Cookies {
		cookies = append(cookies, cookie.Name+"="+cookie.Value)
	}
	request.Headers["cookie"] = strings.Join(cookies, ";")
	return f.fileSpider.Do(request)
	//resp.Header.Add("Content-Type", writer.FormDataContentType())
	//result, err := crawler(url, dataMap["method"].(string), resp.Header, payload)
	//if err != nil {
	//	return Response{}, nil
	//}
	//return Response{
	//	body: []byte(result),
	//}, nil
}

func crawler(url, method string, header http.Header, payload *bytes.Buffer) (string, error) {

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header = header
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), nil
}
