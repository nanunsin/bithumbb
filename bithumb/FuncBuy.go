package bithumb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// Market buy
func (b *Bithumb) BuyETH(units float64) (info [5]Market_Info) {

	var market_json_info market_buy_json_rec
	params := fmt.Sprintf("units=%.1f&currency=ETH", units)

	fmt.Printf("params: %s\n", params)
	resp_data_str := b.apiCall("/trade/market_buy", params)

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
			fmt.Printf(" - Cont_id : %s\n", value.Cont_id)
			fmt.Printf(" - Price : %.2f\n", value.Price)
			fmt.Printf(" - Total : %.2f\n", value.Total)

			info[i].Cont_id = value.Cont_id
			info[i].Price = value.Price
			info[i].Units = value.Units
		}

	} else {
		log.Printf("%s\n", resp_data_str)
	}
	return
}

/*

func (b *Bithumb) BuyETHEx(info *Market_Info, units float64) {

	params := fmt.Sprintf("price=%d&units=%.1f&order_currency=ETH", units)

	var market_json_info market_buy_json_rec

}

*/

func setPlaceTradeParam(price int, unit float64) string {
	v := url.Values{}

	Price := fmt.Sprintf("%d", price)
	Units := fmt.Sprintf("%.1f", unit)

	v.Set("price", Price)
	v.Set("units", Units)

	return v.Encode()
}
