package main

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Parser struct {
	Url string
	Collector *colly.Collector
}

func NewParser(url string) *Parser {
	parser := &Parser{Url: url}
	parser.Collector = colly.NewCollector(colly.AllowURLRevisit())
	return parser
}

func (p *Parser) setup() {
	p.Collector.OnHTML("a.tc-item", func(item *colly.HTMLElement) {
		sSide := item.ChildText("div.tc-side")
		sServer := item.ChildText("div.tc-server")
		sName := item.ChildText("div.tc-user .media-user-name span")
		sReviewsMedian := item.DOM.Find("div.tc-user .media-user-reviews .rating-stars .fas")
		sReviewsCount := item.ChildText("div.tc-user div.media-user-reviews span.rating-mini-count")
		sAmount := item.ChildText("div.tc-amount")
		sPrice := item.ChildText("div.tc-price")

		sAmount = strings.ReplaceAll(sAmount, " ", "")
		sPrice = strings.ReplaceAll(sPrice, " ", "")
		sPrice = strings.ReplaceAll(sPrice, string(rune(8381)), "")

		entry := new(Entry)
		entry.Side = sSide
		entry.Server = sServer
		entry.Seller = new(Seller)
		entry.Seller.Name = sName
		entry.Seller.ReviewsMedian = int64(sReviewsMedian.Length())
		entry.Seller.ReviewsCount, _ = strconv.ParseInt(sReviewsCount, 10, 64)
		entry.Amount, _ = strconv.ParseInt(sAmount, 10, 64)
		entry.Price, _ = strconv.ParseFloat(sPrice, 64)
		entry.Timestamp = time.Now()

		entry.ID = entry.getID()

		p.Proceed(entry)
	})

	p.Collector.OnScraped(func(response *colly.Response) {
		p.Clear()
		if !p.Initialized {
			p.Initialized = true
		}
	})
}


func (p *Parser) Run(wg *sync.WaitGroup) {
	defer wg.Done()



}
