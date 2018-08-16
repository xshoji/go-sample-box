package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/xshoji/go-sample-box/webscraping/client"
)

func main() {
	clientNewsBitcoinFeed := client.NewClient("https://news.bitcoin.com/feed")
	fmt.Println("[Get]")
	rss := clientNewsBitcoinFeed.Get("/").GetBodyAsObject()
	fmt.Println("Title : " + rss.RssChanel.Title)
	fmt.Println("Link : " + rss.RssChanel.Link)
	fmt.Println("Description : " + rss.RssChanel.Description)
	countItems := len(rss.RssChanel.Items)
	for index := 0; index < countItems; index++ {
		fmt.Println("--")
		fmt.Println("  " + rss.RssChanel.Items[index].Guid)
		fmt.Println("  " + rss.RssChanel.Items[index].Title)
		fmt.Println("  " + rss.RssChanel.Items[index].Link)
		doc, _ := goquery.NewDocument(rss.RssChanel.Items[index].Link)
		// - [goqueryでお手軽スクレイピング！ - Qiita](https://qiita.com/yosuke_furukawa/items/5fd41f5bcf53d0a69ca6#goquery%E3%82%92%E4%BD%BF%E3%81%86)
		doc.Find("p").Each(func(_ int, s *goquery.Selection) {
			test := s.Text()
			fmt.Println(test)
		})
		countCategories := len(rss.RssChanel.Items[index].Categories)
		for index2 := 0; index2 < countCategories; index2++ {
			fmt.Println("    " + rss.RssChanel.Items[index].Categories[index2])
		}
	}
}
