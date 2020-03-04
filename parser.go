package main

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"sync"
)

type Parser struct {
	Url string
	Collector *colly.Collector
	Entries Entries
}

func NewParser(url string) *Parser {
	parser := &Parser{Url: url, Entries: Entries{}}
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

		tSide := sSide
		tServer := sServer
		tSellerName := sName
		tSellerReviewsMedian := int64(sReviewsMedian.Length())
		tSellerReviewsCount, _ := strconv.ParseInt(sReviewsCount, 10, 64)
		tAmount, _ := strconv.ParseInt(sAmount, 10, 64)
		tPrice, _ := strconv.ParseFloat(sPrice, 64)

		entry := NewEntry(tServer, tSide, tSellerName, tSellerReviewsCount, tSellerReviewsMedian, tAmount, tPrice)

		p.Proceed(entry)
	})

	p.Collector.OnScraped(func(response *colly.Response) {
		p.Clear()
	})
}

func (p *Parser) Proceed(entry *Entry) {

}

func (p *Parser) Clear() {

}


func (p *Parser) Run(wg *sync.WaitGroup) {
	defer wg.Done()



}
