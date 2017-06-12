package bithumb

import "time"

type Ticker_info struct {
	Date  time.Time
	Price float64
	Min   float64
	Max   float64
}

type Market_Info struct {
	Cont_id string
	Units   float64
	Price   float64
	Total   float64
	Fee     float64
}

type WMP struct {
	Price     float64
	Units     float64
	Ask       int
	AskUnit   float64
	RecentAsk float64
	Bid       int
	BidUnit   float64
	RecentBid float64
}
