package client

// Rss Rss
// - [XMLのパース/生成 - はじめてのGo言語](http://cuto.unirita.co.jp/gostudy/post/standard-library-xml/)
type Rss struct {
	RssChanel RssChannel `xml:"channel"`
}

// RssChannel RssChannel
type RssChannel struct {
	Title           string `xml:"title"`
	Link            string `xml:"link"`
	Description     string `xml:"description"`
	LastBuildDate   string `xml:"lastBuildDate"`
	UpdatePeriod    string `xml:"sy:updatePeriod"`
	UpdateFrequency string `xml:"sy:updateFrequency"`
	Items           []Item `xml:"item"`
}

// Item Item
type Item struct {
	Guid        string   `xml:"guid"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Categories  []string `xml:"category"`
}
