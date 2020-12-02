package EightBit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	PublicApi = "https://public.bitbank.cc/"
)

type CryptoJson struct {
	Success int `json:"success"`
	Data MarketData `json:"data"`
}

type MarketData struct {
	Sell string `json:"sell"`
	Buy string `json:"buy"`
	High string `json:"high"`
	Low string `json:"low"`
	Last string `json:"last"`
	Vol string `json:"vol"`
	Timestamp int64 `json:"timestamp"`
}

func GetCryptoJpy(name string) MarketData {
	req, err := http.NewRequest("GET", PublicApi+ name + "_jpy" + "/ticker", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("EightBit", "OSS")
	client := new(http.Client)

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cj CryptoJson
	err = json.Unmarshal([]byte(body), &cj)
	if err != nil {
		fmt.Println(err)
	}

	if cj.Success != 1 {
		fmt.Println("[EightBit] Failed API request")
	}

	return cj.Data
}

func Get_btc_jpy() MarketData {
	return GetCryptoJpy("btc")
}
func Get_xrp_jpy() MarketData {
	return GetCryptoJpy("xrp")
}
func Get_ltc_jpy() MarketData {
	return GetCryptoJpy("ltc")
}
func Get_eth_jpy() MarketData {
	return GetCryptoJpy("eth")
}
func Get_mona_jpy() MarketData {
	return GetCryptoJpy("mona")
}
func Get_bcc_jpy() MarketData {
	return GetCryptoJpy("bcc")
}
func Get_xlm_jpy() MarketData {
	return GetCryptoJpy("xlm")
}