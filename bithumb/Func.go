package bithumb

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
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

	info.Price = getRightPrice(price / float64(count))

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

<<<<<<< HEAD
	var transJSON recTransactions
	respDataStr := b.publicApiCall("/public/recent_transactions/ETH", "count=10")
	//	fmt.Printf("%s\n", resp_data_str)
=======
	params := fmt.Sprintf("offset=0&count=%d", 100)
	var trans_json_rec_info recTransactions
	resp_data_str := b.publicApiCall("/public/recent_transactions/ETH", params)
	//fmt.Printf("%s\n", resp_data_str)
>>>>>>> 971a5881395769ddb66efd4d4ff8eb61c3840936

	respDataBytes := []byte(respDataStr)

	err := json.Unmarshal(respDataBytes, &transJSON)
	if err != nil {
		log.Println(err.Error())
		return
	}

<<<<<<< HEAD
	var bidSum, askSum float64

	for _, value := range transJSON.Data {
		t := AnalyzeDate(value.Date)
		if b.lasttime.Equal(t) {
			log.Printf("out:%v=%v\n", b.lasttime, t)
			return
		}

		//		if value.Units > 5 {

		fmt.Printf("[%s] %.f\t%.5f\t%s\n", value.Type, value.Price, value.Units, t)
		//		}

		switch value.Type {
		case "bid":
			bidSum += value.Units
		case "ask":
			askSum += value.Units
		}
	}

	fmt.Println("----------------")
	//fmt.Printf("bid: %.5f\nask: %.5f\n", bidSum, askSum)
	b.lasttime = AnalyzeDate(transJSON.Data[0].Date)
=======
	//count := 100
	count := len(trans_json_rec_info.Data)
	fmt.Printf("len : %d\n", count)

	bidSum := 0.0
	askSum := 0.0

	for i := 0; i < count; i++ {
		if trans_json_rec_info.Data[i].Units > 5 {
			t := AnalyzeDate(trans_json_rec_info.Data[i].Date)
			fmt.Printf("[%s] %.f\t%.f\t%s\n", trans_json_rec_info.Data[i].Type, trans_json_rec_info.Data[i].Price, trans_json_rec_info.Data[i].Units, t)

			switch trans_json_rec_info.Data[i].Type {
			case "bid":
				bidSum += trans_json_rec_info.Data[i].Units
			case "ask":
				askSum += trans_json_rec_info.Data[i].Units
			}
		}
	}

	fmt.Println("-------------")
	fmt.Printf("Bid : %.5f\n", bidSum)
	fmt.Printf("Ask : %.5f\n", askSum)
>>>>>>> 971a5881395769ddb66efd4d4ff8eb61c3840936
}

func (b *Bithumb) GetETHOrders(orderbook *OrderBook) bool {

	var orderbookJSON orderbookJson
	resp_data_str := b.apiCall("/public/orderbook/ETH", "")
	//	fmt.Printf("%s\n", resp_data_str)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &orderbookJSON)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if orderbookJSON.Status != "0000" {
		log.Printf("API Call Failed, %s\n", orderbookJSON.Status)
		return false
	}

	for i := 0; i < 10; i++ {
		BidData := orderbookJSON.Data.Bids[i]
		AskData := orderbookJSON.Data.Asks[i]
		// Copy
		orderbook.Bids[i].Price = BidData.Price
		orderbook.Bids[i].Quantity = BidData.Quantity

		orderbook.Asks[i].Price = AskData.Price
		orderbook.Asks[i].Quantity = AskData.Quantity
	}
	return true
}

//////////////////////////////////
// private
func getRightPrice(price float64) (rightPrice float64) {
	upper := math.Floor(price/100.0) * 100.0
	lower := math.Mod(price, 100.0)

	if lower > 50.0 {
		lower = 50.0
	} else {
		lower = 0.0
	}
	rightPrice = upper + lower
	return
}

// Util Function

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

func AnalyzeDate(dateSrc string) time.Time {
	date := strings.Split(dateSrc, " ")

	year, month, day := time.Now().Date()

	hhmmss := strings.Split(date[1], ":")
	hh, _ := strconv.Atoi(hhmmss[0])
	mm, _ := strconv.Atoi(hhmmss[1])
	ss, _ := strconv.Atoi(hhmmss[2])

	local, _ := time.LoadLocation("Local")

	return time.Date(year, month, day, hh, mm, ss, 0, local)
}
