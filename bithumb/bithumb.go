package bithumb

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
	@fn microsectime
	@brief Return current microseconds.
	@return Returns the current time measured in the number of microseconds(Unix timestamp + microseconds) since the Unix Epoch (January 1 1970 00:00:00 GMT).
*/

func microsectime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

/*
	@fn hash_hmac
	@brief Generate a keyed hash value using the HMAC method.
	@param hmac_key: Shared secret key used for generating the HMAC variant of the message digest.
	@param hmac_data: Message to be hashed.
	@return Returns a string containing the calculated message digest as base64 encoded hexits.
*/

func hash_hmac(hmac_key string, hmac_data string) (hash_hmac_str string) {
	hmh := hmac.New(sha512.New, []byte(hmac_key))
	hmh.Write([]byte(hmac_data))

	hex_data := hex.EncodeToString(hmh.Sum(nil))
	hash_hmac_bytes := []byte(hex_data)
	hmh.Reset()

	hash_hmac_str = base64.StdEncoding.EncodeToString(hash_hmac_bytes)

	return (hash_hmac_str)
}

type Bithumb struct {
	api_key    string
	api_secret string
	lasttime   time.Time
}

func NewBithumb(key, secret string) *Bithumb {

	t := time.Now().Round(time.Second)
	return &Bithumb{
		api_key:    key,
		api_secret: secret,
		lasttime:   t,
	}
}

func (b *Bithumb) publicApiCall(endpoint, params string) (resp_data_str string) {
	var api_url = "https://api.bithumb.com"

	params = "?" + params

	// Connects to Bithumb API server and returns JSON result value.
	client := &http.Client{}
	http_req, _ := http.NewRequest("GET", api_url+endpoint+params, nil) // URL-encoded payload

	resp, err := client.Do(http_req)
	if err != nil {
		return ("")
	}

	resp_data, err := ioutil.ReadAll(resp.Body)
	resp_data_str = string(resp_data)

	return (resp_data_str)
}

func (b *Bithumb) apiCall(endpoint, params string) (resp_data_str string) {
	var api_url = "https://api.bithumb.com"

	e_endpoint := url.QueryEscape(endpoint)
	params += "&endpoint=" + e_endpoint

	hmac_key := b.api_secret

	// Api-Nonce information generation.
	nonce_int64 := microsectime()
	api_nonce := fmt.Sprint(nonce_int64)

	// Api-Sign information generation.
	hmac_data := endpoint + string(0) + params + string(0) + api_nonce
	hash_hmac_str := hash_hmac(hmac_key, hmac_data)
	api_sign := hash_hmac_str

	// Connects to Bithumb API server and returns JSON result value.
	client := &http.Client{}
	http_req, _ := http.NewRequest("POST", api_url+endpoint, bytes.NewBufferString(params)) // URL-encoded payload

	http_req.Header.Add("Api-Key", b.api_key)
	http_req.Header.Add("Api-Sign", api_sign)
	http_req.Header.Add("Api-Nonce", api_nonce)
	http_req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	content_length_str := strconv.Itoa(len(params))
	http_req.Header.Add("Content-Length", content_length_str)

	resp, err := client.Do(http_req)
	if err != nil {
		return ("")
	}

	resp_data, err := ioutil.ReadAll(resp.Body)
	resp_data_str = string(resp_data)

	return (resp_data_str)
}

func (b *Bithumb) publicApiCall(endpoint, params string) (resp_data_str string) {
	var api_url = "https://api.bithumb.com"

	params = "?" + params

	// Connects to Bithumb API server and returns JSON result value.
	client := &http.Client{}
	http_req, _ := http.NewRequest("GET", api_url+endpoint+params, nil) // URL-encoded payload

	resp, err := client.Do(http_req)
	if err != nil {
		return ("")
	}

	resp_data, err := ioutil.ReadAll(resp.Body)
	resp_data_str = string(resp_data)

	return (resp_data_str)
}

/*
func (b *Bithumb) GetETHPrice() string {
	fmt.Println("Bithumb Public API URI('/public/ticker') Request...")

	var ticker_json_rec_info ticker_json_rec
	resp_data_str := b.apiCall("/public/ticker/ETH", "")
	//	fmt.Printf("%s\n", resp_data_str)

	resp_data_bytes := []byte(resp_data_str)

	json.Unmarshal(resp_data_bytes, &ticker_json_rec_info)

	return fmt.Sprintf("ETH: %.2f\n", ticker_json_rec_info.Data.Closing_price)
}
*/
