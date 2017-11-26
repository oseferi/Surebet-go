package chrome

import (
	"github.com/knq/chromedp"
	"github.com/knq/chromedp/cdp"
	"io/ioutil"
	"context"
	"log"
	"surebetSearch/common"
)

func SaveScn(url string) chromedp.ActionFunc {
	return func(ctxt context.Context, c cdp.Handler) error {
		var buf []byte
		chromedp.CaptureScreenshot(&buf).Do(ctxt, c)
		siteName, err := common.GetSiteName(url)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(siteName+".png", buf, 0644); err != nil {
			return err
		}
		log.Println("Screenshot saved")
		return nil
	}
}

func GetHtml(html *string) chromedp.ActionFunc {
	return func(ctxt context.Context, c cdp.Handler) error {
		chromedp.OuterHTML("html", html).Do(ctxt, c)
		return nil
	}
}
