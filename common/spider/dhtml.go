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
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var (
		res string
		ids []cdp.NodeID
	)
	if request.ProxyCallBack != nil {
		chromedp.ProxyServer(request.ProxyCallBack())
	}
	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPad),
		chromedp.EmulateViewport(1024, 2048, chromedp.EmulateScale(2)),
		chromedp.Navigate(request.Url),
		chromedp.NodeIDs(`document`, &ids, chromedp.ByJSPath),
		chromedp.Sleep(request.Timeout),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			res, err = dom.GetOuterHTML().WithNodeID(ids[0]).Do(ctx)
			return err
		}),
	); err != nil {
		logs.S.Warnf("chromedp error: %s", err)
		return Response{}, err
	}
	return Response{
		Body:       []byte(res),
		SpiderType: d.spiderType,
	}, nil
}
