package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "data_feed_processor"
)

type Poloniex struct {
	// api keys
	api_key    string
	api_secret string
	client     *http.Client
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

func New(client *http.Client, api_key string, api_secret string) *Poloniex {
	return &Poloniex{"332ZVG1K-ZJKEHYOS-CE9FJWX0-4UYFATS4", "b1b72db287909e1d628d6eb339f8c1788aeed2425dee2d90e75b9ba81ebf7af9a120a831b72338a942619effb96ebb4c3a67782c2c4071394940d35e044ad45f", &http.Client{}}
}
func main() {
	//configure postgresql database
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Error")
	}

	defer db.Close()

	user1 := Poloniex{
		api_key:    "332ZVG1K-ZJKEHYOS-CE9FJWX0-4UYFATS4",
		api_secret: "b1b72db287909e1d628d6eb339f8c1788aeed2425dee2d90e75b9ba81ebf7af9a120a831b72338a942619effb96ebb4c3a67782c2c4071394940d35e044ad45f",
		client:     &http.Client{},
	}
	user1.getPoloniexData()

}

func (p *Poloniex) getPoloniexData() string {

	// Build the request
	url := "https://poloniex.com/public?command=returnTradeHistory&currencyPair=BTC_NXT&start=1410158341&end=1410499372"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	res, _ := p.client.Do(req)

	fmt.Println(res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data poloniexData
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data.Types)
	os.Exit(0)
	return "Okay"
}
