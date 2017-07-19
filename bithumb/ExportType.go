package bithumb

import "time"

type Ticker_info struct {
	Date  time.Time
	Price float64
	Min   float64
	Max   float64
}

type Market_Info struct {
	ContID  string
	OrderID string
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

type OrderBook struct {
	Bids [10]OrderDetail
	Asks [10]OrderDetail
}

type OrderDetail struct {
	Price    uint64
	Quantity float64
}

// trade/place
type PlaceInfo struct {
	ContID string
	Units  float64
	Price  float64
	Total  float64
	Fee    float64
}
