package bithumb

import (
	"encoding/json"
	"fmt"
	"log"
)

func (b *Bithumb) SellETH(units float64) (info [5]Market_Info) {

	var market_json_info market_buy_json_rec
	params := fmt.Sprintf("units=%.1f&currency=ETH", units)

	fmt.Printf("params: %s\n", params)
	resp_data_str := b.apiCall("/trade/market_sell", params)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &market_json_info)
	if err != nil {
		log.Printf("%s\n", resp_data_str)
		log.Println(err.Error())
		return
	}

	if market_json_info.Status == "0000" {
		fmt.Printf(" - Order_id : %s\n", market_json_info.Order_id)
		for i, value := range market_json_info.Data {
			fmt.Printf("[%d]\n", i)
			fmt.Printf(" - ContID : %s\n", value.ContID)
			fmt.Printf(" - Price : %.2f\n", value.Price)
			fmt.Printf(" - Total : %.2f\n", value.Total)

			info[i].ContID = value.ContID
			info[i].Price = value.Price
			info[i].Units = value.Units
		}

	} else {
		log.Printf("%s\n", resp_data_str)
	}
	return
}

func (b *Bithumb) SellPlaceETH(price int, units float64) (info [5]Market_Info) {

	var market_json_info market_buy_json_rec
	params := fmt.Sprintf("price=%d&units=%.1f&order_currency=ETH&type=ask", price, units)

	fmt.Printf("params: %s\n", params)
	resp_data_str := b.apiCall("/trade/place", params)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &market_json_info)
	if err != nil {
		log.Printf("%s\n", resp_data_str)
		log.Println(err.Error())
		return
	}

	if market_json_info.Status == "0000" {
		fmt.Printf(" - Order_id : %s\n", market_json_info.Order_id)
		for i, value := range market_json_info.Data {
			fmt.Printf("[%d]\n", i)
			fmt.Printf(" - ContID : %s\n", value.ContID)
			fmt.Printf(" - Price : %.2f\n", value.Price)
			fmt.Printf(" - Total : %.2f\n", value.Total)

			info[i].ContID = value.ContID
			info[i].Price = value.Price
			info[i].Units = value.Units
		}

	} else {
		log.Printf("%s\n", resp_data_str)
	}
	return
}

func (b *Bithumb) SellETHEx(units float64) (info [5]Market_Info) {

	var market_json_info market_buy_json_rec
	params := fmt.Sprintf("units=%.1f&currency=ETH", units)

	fmt.Printf("params: %s\n", params)
	resp_data_str := b.apiCall("/trade/market_sell", params)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &market_json_info)
	if err != nil {
		log.Printf("%s\n", resp_data_str)
		log.Println(err.Error())
		return
	}

	if market_json_info.Status == "0000" {
		fmt.Printf(" - Order_id : %s\n", market_json_info.Order_id)
		for i, value := range market_json_info.Data {
			fmt.Printf("[%d]\n", i)
			fmt.Printf(" - ContID : %s\n", value.ContID)
			fmt.Printf(" - Price : %.2f\n", value.Price)
			fmt.Printf(" - Total : %.2f\n", value.Total)

			info[i].ContID = value.ContID
			info[i].Price = value.Price
			info[i].Units = value.Units
		}

	} else {
		log.Printf("%s\n", resp_data_str)
	}
	return
}

func (b *Bithumb) SellPlace(coin string, price int, units float64) (info [5]Market_Info) {

	var market_json_info market_buy_json_rec
	params := fmt.Sprintf("price=%d&units=%.1f&order_currency=%s&type=ask", price, units, coin)

	fmt.Printf("params: %s\n", params)
	resp_data_str := b.apiCall("/trade/place", params)

	resp_data_bytes := []byte(resp_data_str)

	err := json.Unmarshal(resp_data_bytes, &market_json_info)
	if err != nil {
		log.Printf("%s\n", resp_data_str)
		log.Println(err.Error())
		return
	}

	if market_json_info.Status == "0000" {
		fmt.Printf(" - Order_id : %s\n", market_json_info.Order_id)
		for i, value := range market_json_info.Data {
			fmt.Printf("[%d]\n", i)
			fmt.Printf(" - ContID : %s\n", value.ContID)
			fmt.Printf(" - Price : %.2f\n", value.Price)
			fmt.Printf(" - Total : %.2f\n", value.Total)

			info[i].ContID = value.ContID
			info[i].Price = value.Price
			info[i].Units = value.Units
		}

	} else {
		log.Printf("%s\n", resp_data_str)
	}
	return
}
