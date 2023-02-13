package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func scraper(url string, pagecount int, saveAs string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	chromedp.ListenTarget(
		ctx,
		func(ev interface{}) {
			if ev, ok := ev.(*network.EventResponseReceived); ok {

				if ev.Type != "XHR" {
					return
				}

				if strings.Contains(ev.Response.URL, "tournament-archive") == false {
					return
				}

				go func() {
					c := chromedp.FromContext(ctx)
					rbp := network.GetResponseBody(ev.RequestID)
					body, err := rbp.Do(cdp.WithExecutor(ctx, c.Target))
					if err != nil {
						fmt.Println(err)
						pagecount--
						fmt.Println("RUN IT AGAIN..")
					}
					if err == nil {
						filename := saveAs + fmt.Sprintf("%02d", pagecount) + ".json"
						err := os.WriteFile(filename, body, 0644)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Printf("SAVED: %v \n", filename)
						fmt.Println("waiting..")
					}
				}()

			}
		},
	)

	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(url+fmt.Sprint(pagecount)),
		chromedp.Sleep(time.Second*time.Duration((10+rand.Intn(15)))),
	)
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	rand.Seed(time.Now().UnixNano())

	var url string
	var pagecount int
	var saveAs string

	flag.StringVar(&url, "u", "https://www.oddsportal.com/baseball/usa/mlb-2022/results/#/page/", "URL to scrap. For now must end as /page/")
	flag.IntVar(&pagecount, "p", 2, "Pages to scrape. MLB 2022 Season has 55 Pages")
	flag.StringVar(&saveAs, "s", "MLB2022-", "Filename for saving")
	flag.Parse()
	fmt.Println("Starting...")
	for i := 1; i <= pagecount; i++ {
		fmt.Printf("CYCLE: %v.. TARGET: %v \n", i, url+fmt.Sprintf("%v", i))
		scraper(url, i, saveAs)
	}
}
