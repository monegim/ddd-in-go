package ch03

import (
	"time"

	"github.com/Rhymond/go-money"
)

type Auction struct {
	ID            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}
