package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

type Seller struct {
	Name string
	ReviewsCount int64
	ReviewsMedian int64
}

type Entry struct {
	ID string
	Server string
	Side string
	Seller *Seller
	Amount int64
	Price float64
	Timestamp time.Time
}

func NewEntry(
		Server string,
		Side string,
		SellerName string,
		SellerReviewsCount int64,
		SellerReviewsMedian int64,
		Amount int64,
		Price float64,
	) *Entry {
	return &Entry{
		ID:        generateID(Server, Side, SellerName),
		Server:    Server,
		Side:      Side,
		Seller:    &Seller{
			Name:          SellerName,
			ReviewsCount:  SellerReviewsCount,
			ReviewsMedian: SellerReviewsMedian,
		},
		Amount:    Amount,
		Price:     Price,
		Timestamp: time.Now(),
	}
}

func generateID(Server, Side, Name string) string {
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%s-%s", Server, Side, Name)))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

type Entries []*Entry
