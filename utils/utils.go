package utils

import (
	"encoding/json"
	"net/http"

	"github.com/OuroborosTFB/Bitcoinbot/model"
)

func GetApiCall() (*model.Price, error) {
	resp, err := http.Get("https://bitex.la/api-v1/rest/btc_usd/market/ticker")
	p := &model.Price{}

	if err != nil {
		return p, err
	}

	err = json.NewDecoder(resp.Body).Decode(p)
	return p, err
}
