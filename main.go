package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const (
	db_user     = "postgres"
	db_password = ""
	db_name     = "data_feed_processor"
)

type Poloniex struct {
	client             *http.Client
	poloniexParameters struct {
		url          string
		currencyPair string
		start        string
		end          string
	}
}

type poloniexData struct {
	GlobalTradeID int64   `json:"globalTradeID"`
	TradeID       string  `json:"tradeID"`
	Date          string  `json:"date"`
	Types         string  `json:"type"`
	Rate          float64 `json:"rate"`
	Amount        float64 `json:"amount"`
	Total         float64 `json:"total"`
}
type Bittrex struct {
	client            *http.Client
	bittrexParameters struct {
		url    string
		market string
	}
}
type bittrexData struct {
	Success string `json:"success"`
	Message string `json:"message"`

	Result []ResultArray `json:"result"`
}

//ResultArray Export the values to ResultArray struct
type ResultArray struct {
	ID        int64   `json:"Id"`
	Timestamp string  `json:"TimeStamp"`
	Quantity  float64 `json:"Quantity"`
	Price     float64 `json:"Price"`
	Total     float64 `json:"Total"`
	Filltype  string  `json:"FillType"`
	Ordertype string  `json:"OrderType"`
}

type databaseInfo struct {
	db_user     string
	db_password string
	db_name     string
}

func main() {

	getHistoricData(1, "BTC-DCR", "", "")

}

func (p *Poloniex) getPoloniexData(currencyPair string, start string, end string) string {

	url := "https://poloniex.com/public?command=returnTradeHistory"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	q := req.URL.Query()
	q.Add("currencyPair", currencyPair)
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	res, _ := p.client.Do(request)

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data poloniexData
	json.Unmarshal(body, &data)
	fmt.Println(string(body))
	fmt.Printf("Results: %v\n", data.Types)
	os.Exit(0)
	return "Okay"
}

func (b *Bittrex) getBittrexData(currencyPair string) {

	url := "https://bittrex.com/api/v1.1/public/getmarkethistory"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	q := req.URL.Query()
	q.Add("market", currencyPair)

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String()) //to test if the url sent is correct

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	res, _ := b.client.Do(request)

	fmt.Println(res.StatusCode) // To check the status code of response

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data bittrexData

	json.Unmarshal(body, &data) // To decode the json data
	// fmt.Println(string(body))
	fmt.Printf("Results: %v\n", data)
	// Loop over structs and display them.
	// for l := range data {
	// 	fmt.Printf("Id = %v, Name = %v", data[l].Message, data[l].Timestamp)
	// 	fmt.Println()
	// }
	return

}

func getHistoricData(exchangeID int, currencyPair string, startTime string, endTime string) {

	if exchangeID == 0 { //poloneix exchange
		user1 := Poloniex{

			client: &http.Client{},
		}
		user1.getPoloniexData(currencyPair, startTime, endTime)

	}

	if exchangeID == 1 { //bittrex Exchange

		user2 := Bittrex{
			client: &http.Client{},
		}
		user2.getBittrexData(currencyPair)
	}
}

// 0 for poloneix
// 1 for bittrex
