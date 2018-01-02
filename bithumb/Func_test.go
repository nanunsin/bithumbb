package bithumb

import (
	"fmt"
	"math"
	"testing"
	"time"

	color "github.com/fatih/color"
)

func TestGetPrice(t *testing.T) {
	t.Log("start")
	bot := NewBithumb("test", "secret")
	var info TickerInfo
	err := bot.GetPrice("ETH", &info)
	if err != nil {
		t.Errorf("err : %v\n", err.Error())
	}

	fmt.Printf("Price: %d\n", info.Price)
}

func TestGetETHPrice(t *testing.T) {

	t.Log("start")
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

func TestGetETHPrice2(t *testing.T) {

	t.Log("start2")
	bot := NewBithumb("test", "secret")
	var info WMP
	err := bot.GetETHPrice(&info)
	if err != nil {
		t.Errorf("err : %v\n", err.Error())
	}

	fmt.Printf("Price: %.2f\n", info.Price)
	fmt.Printf("Gap: (%.2f - %.2f) = %.2f\n", info.RecentBid, info.RecentAsk, info.RecentBid-info.RecentAsk)

	upper := math.Floor(info.Price/100.0) * 100.0
	lower := math.Mod(info.Price, 100.0)

	if lower > 50.0 {
		lower = 50.0
	} else {
		lower = 0.0
	}

	fmt.Printf("Right Price : %.f\n", upper+lower)
}

func TestGetETHOrders(t *testing.T) {

	t.Log("start3")
	bot := NewBithumb("test", "secret")
	var info OrderBook
	result := bot.GetETHOrders(&info)
	if !result {
		t.Errorf("err\n")
	}

	for i := 9; i >= 0; i-- {
		data := info.Asks[i]
		c := color.New(color.FgRed)
		c.Printf("%d | %.6f \n", data.Price, data.Quantity)
	}

	for i := 0; i < 10; i++ {
		data := info.Bids[i]
		c := color.New(color.FgGreen)
		c.Printf("%d | %.6f \n", data.Price, data.Quantity)
	}
}

func TestGetETHTrans(t *testing.T) {

	t.Log("start3")
	bot := NewBithumb("test", "secret")
	bot.GetETHRecTrans()
	time.Sleep(3)
	bot.GetETHRecTrans()
	time.Sleep(3)
	bot.GetETHRecTrans()
}
