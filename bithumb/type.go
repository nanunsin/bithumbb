package bithumb

// /public/ticker structure

type ticker_rec struct {
	Opening_price float64 `json:"opening_price,string"`
	Closing_price float64 `json:"closing_price,string"`
	Min_price     float64 `json:"min_price,string"`
	Max_price     float64 `json:"max_price,string"`
	Date          int64   `json:"date,string"`
}

// Ticker JSON structure

type ticker_json_rec struct {
	Status string     `json:"status"`
	Data   ticker_rec `json:"data"`
}

// Recent_transactions structure
type trans_rec struct {
	Date  string  `json:"transaction_date"`
	Type  string  `json:"type"`
	Units float64 `json:"units_traded,string"`
	Price float64 `json:"price,string"`
	Total int64   `json:"total,string"`
}

// Ticker JSON structure

type tickerInfo struct {
	ClosingPrice int64 `json:"closing_price,string"`
	Date         int64 `json:"date,string"`
}

type tickerJson struct {
	Status string     `json:"status"`
	Data   tickerInfo `json:"data"`
}

// Recent_transactions JSON structure
type trans_json_rec struct {
	Status string      `json:"status"`
	Data   []trans_rec `json:"data"`
}

// /info/account structure

type account_rec struct {
	Created    int64   `json:"created,string"`
	Account_id string  `json:"account_id"`
	Trade_fee  float64 `json:"trade_fee,string"`
	Balance    float64 `json:"balance,string"`
}

// Account JSON structure

type account_json_rec struct {
	Status string      `json:"status"`
	Data   account_rec `json:"data"`
}

// /trade/market_buy struct

type market_buy_json_rec struct {
	Status   string       `json:"status"`
	Order_id string       `json:"order_id"`
	Data     []market_rec `json:"data"`
}

type market_rec struct {
	ContID  string  `json:"cont_id"`
	OrderID string  `json:"order_id"`
	Units   float64 `json:"units,string"`
	Price   float64 `json:"price,string"`
	Total   float64 `json:"total"`
	Fee     float64 `json:"fee"`
}

// /public/orderbook struct
type orderbookJson struct {
	Status string        `json:"status"`
	Data   orderbookData `json:"data"`
}
type orderbookData struct {
	Timestamp     uint64           `json:"timestamp,string"`
	OrderCurrency string           `json:"order_currency"`
	Bids          []orderbookPrice `json:"bids"`
	Asks          []orderbookPrice `json:"asks"`
}

type orderbookPrice struct {
	Quantity float64 `json:"quantity,string"`
	Price    uint64  `json:"price,string"`
}

//  /trade/place

type placeJson struct {
	Status  string      `json:"status"`
	OrderID string      `json:"order_id"`
	Data    []placeData `json:"data"`
}

type placeData struct {
	ContID string  `json:"cont_id"`
	Units  float64 `json:"units,string"`
	Price  float64 `json:"price,string"`
	Total  float64 `json:"total"`
	Fee    float64 `json:"fee"`
}

// /public/recent_transactions
type transactiondata struct {
	Date  string  `json:"transaction_date"`
	Type  string  `json:"type"`
	Units float64 `json:"units_traded,string"`
	Price float64 `json:"price,string"`
	Total float64 `json:"total,string"`
}

type recTransactions struct {
	Status string            `json:"status"`
	Data   []transactiondata `json:"data"`
}

type DefaultReturn struct {
	Status string `json:"status"`
}

// /private /info/balance
type balancedata struct {
	TotalBTC  float64 `json:"total_btc,string"`
	TotalETH  float64 `json:"total_eth,string"`
	TotalLTC  float64 `json:"total_ltc,string"`
	TotalETC  float64 `json:"total_etc,string"`
	TotalEOS  float64 `json:"total_eos,string"`
	TotalXRP  float64 `json:"total_xrp,string"`
	TotalQTUM float64 `json:"total_qtum,string"`
	TotalKRW  int64   `json:"total_krw"`
}

type balanceJson struct {
	Status string      `json:"status"`
	Data   balancedata `json:"data"`
}

// /private /info/order_detail

type orderdetaildata struct {
	Type          string  `json:"type"`
	OrderCurrency string  `json:"order_currency"`
	Units         float64 `json:"units_traded,string"`
	Price         int64   `json:"price,string"`
}

type orderdetailJson struct {
	Status string            `json:"status"`
	Data   []orderdetaildata `json:"data"`
}
