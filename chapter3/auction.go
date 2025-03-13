package chapter3

import (
	"github.com/Rhymond/go-money"
	"time"
)

type Auction struct {
	ID            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
