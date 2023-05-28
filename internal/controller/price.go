package controller

import (
	"bitcoin-notify/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinGeckoResponse struct {
	Bitcoin map[string]float64 `json:"bitcoin"`
}

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	url := config.GetBitcoinPriceApi
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var data CoinGeckoResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	price := data.Bitcoin["uah"]

	fmt.Println("Bitcoin Price (UAH):", price)
}
