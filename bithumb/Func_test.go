package bithumb

import (
	"fmt"
	"testing"
)

func TestGetETHPrice(t *testing.T) {
	t.Log("test1")

	bot := NewBithumb("test", "secret")
	var info WMP
	err := bot.GetETHPrice(&info)
	if err != nil {
		t.Errorf("err : %v\n", err.Error())
	}

	fmt.Printf("Price: %.2f\n", info.Price)
	fmt.Printf("Bid: %d (%.2f)\n", info.Bid, info.BidUnit)
	fmt.Printf("Ask: %d (%.2f)\n", info.Ask, info.AskUnit)
	if info.RecentAsk != 0.0 && info.RecentBid != 0.0 {
		fmt.Printf("Gap: (%.2f - %.2f) = %.2f\n", info.RecentBid, info.RecentAsk, info.RecentBid-info.RecentAsk)
	}
}
