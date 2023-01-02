package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/jessevdk/go-flags"
	"net/url"
	"os"
	"regexp"
	"strings"
	"github.com/yuin/charsetutil"
	"sync"
)

type options struct {
	Url string `short:"u" long:"url" description:"A target url" required:"true"`
	ConcurrentLinkScraping int `short:"c" long:"concurrentLinkScraping" description:"A number of concurrent for scraping of link url" default:"3"`
}

func main() {

	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		// some error handling
		return
	}
	concurrentLinkScraping := opts.ConcurrentLinkScraping
	url, err := url.ParseRequestURI(opts.Url)
	if err != nil {
		fmt.Printf("[ERROR] %v is invalid url format.\n", opts.Url)
		os.Exit(1)
	}
	validateScheme(url.Scheme)
	validateHost(url.Host)
	fmt.Printf("Info\tScheme\t%v\n", url.Scheme)
	fmt.Printf("Info\tHost\t%v\n", url.Host)
	fmt.Printf("Info\tPath\t%v\n", url.Path)
	// > Go 言語で改行コードを変換する（正規表現以外の解） - Qiita
	// > https://qiita.com/spiegel-im-spiegel/items/f1cc014ecb233afaa8af
	replacer := strings.NewReplacer(
		"\r\n", " ",
		"\r", " ",
		"\n", " ",
		"\t", " ",
	)
	replacerSpace := strings.NewReplacer(
		" ", "",
	)

	doc, _ := goquery.NewDocument(url.String())
	//----------------
	// Title, Description, Keywords取得
	//----------------
	// - [goqueryでお手軽スクレイピング！ - Qiita](https://qiita.com/yosuke_furukawa/items/5fd41f5bcf53d0a69ca6#goquery%E3%82%92%E4%BD%BF%E3%81%86)
	doc.Find("title").Each(func(_ int, s *goquery.Selection) {
		fmt.Printf("Header\tTitle\t%v\n", s.Text())
	})
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		attribute, _ := s.Attr("name")
		content, _ := s.Attr("content")
		if attribute == "description" {
			fmt.Printf("Header\tDescription\t%v\n", content)
		}
		if attribute == "keywords" {
			fmt.Printf("Header\tkeywords\t%v\n", content)
		}

		attribute, _ = s.Attr("property")
		if attribute == "og:site_name" {
			fmt.Printf("Header\tOpenGraphSiteName\t%v\n", content)
		}
		if attribute == "og:title" {
			fmt.Printf("Header\tOpenGraphTitle\t%v\n", content)
		}
		if attribute == "og:description" {
			fmt.Printf("Header\tOpenGraphDescription\t%v\n", content)
		}
		if attribute == "og:url" {
			fmt.Printf("Header\tOpenGraphUrl\t%v\n", content)
		}
		if attribute == "og:image" {
			fmt.Printf("Header\tOpenGraphImage\t%v\n", content)
		}
	})
	//----------------
	// Imageリンク取得
	//----------------
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		imageUrl, _ := s.Attr("src")
		imageTitle, _ := s.Attr("alt")
		if imageUrl == "" {
			return
		}
		// 相対パスの場合は絶対パスへ変換する
		if checkRegexp(`^/`, imageUrl) == true {
			imageUrl = url.Scheme + "://" + url.Host + imageUrl
		}
		fmt.Printf("Body\tImageUrl\t%v\t%v\n", strings.TrimSpace(imageTitle), imageUrl)
	})
	//----------------
	// アンカーリンク取得
	//----------------
	channel := make(chan int, concurrentLinkScraping)
	waitGroup := sync.WaitGroup{}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		channel <- 1
		waitGroup.Add(1)
		go func() {
			// > go - golang sync.WaitGroup never completes - Stack Overflow
			// > https://stackoverflow.com/questions/27893304/golang-sync-waitgroup-never-completes
			defer func() {
				<- channel
				waitGroup.Done()
			}()
			linkUrl, _ := s.Attr("href")
			if linkUrl == "" {
				return
			}
			// 相対パスの場合は絶対パスへ変換する
			if checkRegexp(`^/`, linkUrl) == true {
				linkUrl = url.Scheme + "://" + url.Host + linkUrl
			}
			// http形式じゃない場合はスルーする
			if checkRegexp(`^http`, linkUrl) == false {
				return
			}

			ankerText := replacer.Replace(s.Text())
			ankerTextNoSpace := replacerSpace.Replace(ankerText)
			// アンカーテキストがない場合はリンク先のタイトルで補完する
			if ankerTextNoSpace == "" {
				doc2, _ := goquery.NewDocument(linkUrl)
				if doc2 != nil {
					// > inforno ：： Goで文字コードを手軽に変換するライブラリ作った
					// > http://inforno.net/articles/2016/07/14/go-charsetutil
					ankerTextByte, _ := charsetutil.EncodeString(doc2.Find("title").Text(), "UTF-8")
					ankerText = string(ankerTextByte)
				}
			}
			// > go - How to trim leading and trailing white spaces of a string? - Stack Overflow
			// > https://stackoverflow.com/questions/22688010/how-to-trim-leading-and-trailing-white-spaces-of-a-string
			fmt.Printf("Body\tLinkUrl\t%v\t%v\n", strings.TrimSpace(ankerText), linkUrl)
		}()
	})
	waitGroup.Wait()
	close(channel)
}

func validateScheme(scheme string) bool {
	// http or httpsじゃない場合はエラー
	if scheme != "http" && scheme != "https" {
		fmt.Printf("[ERROR] %v is invalid scheme.\n", scheme)
		os.Exit(1)
	}
	return true
}

func validateHost(host string) bool {
	// xxx.xxx形式じゃない場合はエラー
	if checkRegexp(`([0-9a-zA-Z]+)\.([0-9a-zA-Z]+)$`, host) == false {
		fmt.Printf("[ERROR] %v is invalid domain format.\n", host)
		os.Exit(1)
	}
	return true
}

// > 逆引きGolang (正規表現)
// > https://ashitani.jp/golangtips/tips_regexp.html
func checkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
