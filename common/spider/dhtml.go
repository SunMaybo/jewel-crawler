package spider

import (
	"context"
	"github.com/SunMaybo/jewel-crawler/logs"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

type DhtmlSpider struct {
	spiderType SpiderType
}

func NewDhtmlSpider() *DhtmlSpider {
	return &DhtmlSpider{
		spiderType: Shtml,
	}
}
func (d *DhtmlSpider) Do(request Request) (Response, error) {
	contextTimeout, timeoutfunc := context.WithTimeout(context.Background(), request.Timeout)
	defer timeoutfunc()
	ctx, cancel := chromedp.NewContext(contextTimeout)
	defer cancel()
	var (
		res string
		ids []cdp.NodeID
	)
	if request.ProxyCallBack != nil {
		proxy := request.ProxyCallBack()
		if proxy != "" {
			chromedp.ProxyServer(proxy)
		}

	}
	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPad),
		chromedp.EmulateViewport(1024, 2048, chromedp.EmulateScale(2)),
		chromedp.Navigate(request.Url),
		chromedp.Sleep(request.Timeout),
		chromedp.NodeIDs(`document`, &ids, chromedp.ByJSPath),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			res, err = dom.GetOuterHTML().WithNodeID(ids[0]).Do(ctx)
			return err
		}),
	); err != nil {
		logs.S.Warnw("chromedp error: %s", err)
		return Response{}, err
	}
	return Response{
		body:       []byte(res),
		SpiderType: d.spiderType,
	}, nil
}
