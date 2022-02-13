package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/OuroborosTFB/Bitcoinbot/models"
)

type Result struct {
	Price models.Price `json:"USD"`
}

func GetPriceAPI() (*models.Price, error) {
	resp, err := http.Get("https://blockchain.info/ticker")
	price := &models.Price{}

	if err != nil {
		return price, err
	}
	body, err := ioutil.ReadAll(resp.Body)

	var result Result
	json.Unmarshal(body, &result)

	price = &result.Price

	return price, err
}
