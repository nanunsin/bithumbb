package bithumb

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"
)

func (b *Bithumb) GetETHPrice(info *WMP) error {

	var trans_json_rec_info trans_json_rec
	resp_data_str := b.apiCall("/public/recent_transactions/ETH", "")
	//	fmt.Printf("%s\n", resp_data_str)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &trans_json_rec_info)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	count := 20
	price := 0.0

	askset := false
	bidset := false

	for i := 0; i < count; i++ {
		price += trans_json_rec_info.Data[i].Price
		info.Units += trans_json_rec_info.Data[i].Units
		switch trans_json_rec_info.Data[i].Type {
		case "bid":
			{
				info.Bid++
				info.BidUnit += trans_json_rec_info.Data[i].Units
				if !bidset {
					info.RecentBid = trans_json_rec_info.Data[i].Price
					bidset = true
				}
			}
		case "ask":
			{
				info.Ask++
				info.AskUnit += trans_json_rec_info.Data[i].Units
				if !askset {
					info.RecentAsk = trans_json_rec_info.Data[i].Price
					askset = true
				}
			}
		}
	}

	info.Price = price / float64(count)

	return nil
}

func (b *Bithumb) GetBTCPrice(info *Ticker_info) error {

	var ticker_json_rec_info ticker_json_rec
	resp_data_str := b.apiCall("/public/ticker/BTC", "")
	//	fmt.Printf("%s\n", resp_data_str)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &ticker_json_rec_info)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	info.Date = time.Unix(ticker_json_rec_info.Data.Date/1000, 0)
	info.Price = ticker_json_rec_info.Data.Closing_price
	info.Min = ticker_json_rec_info.Data.Min_price
	info.Max = ticker_json_rec_info.Data.Max_price

	return nil
}

//////////////////////////////////
// public

func (b *Bithumb) GetETHRecTrans() {

	var trans_json_rec_info trans_json_rec
	resp_data_str := b.apiCall("/public/recent_transactions/ETH", "")
	//	fmt.Printf("%s\n", resp_data_str)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &trans_json_rec_info)
	if err != nil {
		log.Println(err.Error())
		return
	}

	count := 20

	for i := 0; i < count; i++ {
		if trans_json_rec_info.Data[i].Units > 5 {
			fmt.Printf("[%s] %.f\t%.f\t%s\n", trans_json_rec_info.Data[i].Type, trans_json_rec_info.Data[i].Price, trans_json_rec_info.Data[i].Units, trans_json_rec_info.Data[i].Date)
		}

	}
}

func GetRightPrice() (price string) {
	bot := NewBithumb("test", "secret")
	var info WMP
	err := bot.GetETHPrice(&info)
	if err != nil {
		log.Printf("[GRP][Error] : %s\n", err.Error())
	}

	upper := math.Floor(info.Price/100.0) * 100.0
	lower := math.Mod(info.Price, 100.0)

	if lower > 50.0 {
		lower = 50.0
	} else {
		lower = 0.0
	}
	price = fmt.Sprintf("%.f", upper+lower)
	return
}
