package spider

import (
	"bytes"
	"errors"
	"github.com/SunMaybo/jewel-crawler/common/parser"
	"mime/multipart"
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
	shtml := NewShtmlSpider(1024)
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
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	if dataMap["params"] != nil {
		paramInters := dataMap["params"].([]interface{})
		for _, paramInter := range paramInters {
			paramMap := paramInter.(map[string]interface{})
			writer.WriteField(paramMap["name"].(string), paramMap["value"].(string))
		}
	}
	err = writer.Close()
	if err != nil {
		return Response{}, nil
	}
	request.Param = payload.String()
	request.Headers["Content-Type"] = writer.FormDataContentType()
	request.Url = url
	request.Method = dataMap["method"].(string)
	return f.fileSpider.Do(request)
}
