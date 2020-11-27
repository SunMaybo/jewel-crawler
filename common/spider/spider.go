package spider

import (
	"github.com/SunMaybo/jewel-crawler/common/spider/charset"
	"github.com/SunMaybo/jewel-crawler/logs"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type SpiderType int

const (
	Shtml SpiderType = iota + 1
	Dhtml
	File
	Api
)

type ProxyCallBack func() string
type SocketProxyCallBack func() (string, string, string)

type CookieCallBack func() []*http.Cookie

type Request struct {
	Url                 string
	Method              string
	Param               string
	Headers             map[string]string
	Timeout             time.Duration
	Retry               int
	ProxyCallBack       ProxyCallBack
	SocketProxyCallBack SocketProxyCallBack
	CookieCallBack      CookieCallBack
}
type Response struct {
	body        []byte
	charset     string
	RedirectUrl string
	Cookies     []*http.Cookie
	Header      http.Header
	SpiderType  SpiderType
}

func (r *Response) GetBytes() []byte {
	return r.body
}
func (r *Response) GetContent() string {
	return string(r.body)
}
func (r *Response) GetCharset() string {
	if r.charset == "" {
		r.charset = r.recognitionCharsetFormat()
	}
	return r.charset
}

// 识别编码方式
func (r *Response) recognitionCharsetFormat() string {
	charsetStr := "utf-8"
	if r.SpiderType == Shtml {
		reg := regexp.MustCompile(`<meta.*?charset=\s*"?(.*?)["|;].*?>`) // 进行正则编译
		findResult := reg.FindStringSubmatch(string(r.body))
		if findResult != nil && len(findResult) > 1 && findResult[1] != "" && strings.Index("gbk,gb18030,gb2312,utf8,utf-8,ansi,big5,unicode,ascii", strings.ToLower(findResult[1])) != -1 {
			charsetStr = findResult[1]
			logs.S.Info("find article code is ------>", charsetStr)
		} else {
			cs, err := charset.GuessBytes(r.GetBytes())
			if err != nil {
				logs.S.Warn("charsetutil cannot find article encode  --------->", err.Error())
			} else {
				charsetStr = cs.Charset()
			}
		}
	}
	logs.S.Infow("article encoding ----->", "charset", charsetStr)
	return charsetStr
}

type Spider interface {
	Do(request Request) (Response, error)
}
