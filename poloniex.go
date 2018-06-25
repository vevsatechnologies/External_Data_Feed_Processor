package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structure containing Poloniex client data

type Poloniex struct {
	client *http.Client
}

//Structure containing Poloniex Historic Data

type poloniexData struct {
	Result []struct {
		GlobalTradeID int64   `json:"globalTradeID"`
		TradeID       string  `json:"tradeID"`
		Date          string  `json:"date"`
		Types         string  `json:"type"`
		Rate          float64 `json:"rate"`
		Amount        float64 `json:"amount"`
		Total         float64 `json:"total"`
	}
}

// Structure containing Poloniex Chart Data

type chartData struct {
	Result []struct {
		Date            int64   `json:"date"`
		High            float64 `json:"high"`
		Low             float64 `json:"low"`
		Open            float64 `json:"open"`
		Close           float64 `json:"close"`
		Volume          float64 `json:"volume"`
		QuoteVolume     float64 `json:"quoteVolume"`
		WeightedAverage float64 `json:"weightedAverage"`
	}
}

// Get Poloniex Historic Data
// parameters : Currency pair, Start time , End time

func (p *Poloniex) getPoloniexData(currencyPair string, start string, end string) string {

	//Initiate Connection with Database

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Error")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//Get Url of Poloniex API from constants.go

	url := poloniexBaseURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	//Append user provided parameters in the URL

	q := req.URL.Query()
	q.Add("command", "returnTradeHistory")
	q.Add("currencyPair", currencyPair)
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	//Get Historic Data from the API

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	res, _ := p.client.Do(request)

	//Get response of the request as []byte

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	//Map the data to poloniexData struct and unmarshal the contents

	var data poloniexData
	json.Unmarshal(body, &data)

	fmt.Printf("Results: %v\n", data)

	//Loop over the entire list to insert data into the table

	for i := range data.Result {
		sqlStatement := `INSERT INTO poloniex_historic_data(globaltradeid,tradeid,date,type,rate,amount,total) VALUES($1,$2,$3,$4,$5,$6,$7)`
		_, err = db.Exec(sqlStatement, data.Result[i].GlobalTradeID, data.Result[i].TradeID, data.Result[i].Date, data.Result[i].Types, data.Result[i].Rate, data.Result[i].Amount, data.Result[i].Total)

	}
	return "Saved poloneix historic data!"
}

//Returns data from Poloniex exchange
//Parameters : currency pair , start time , end time

func (p *Poloniex) fetchPoloniexData(date string) {

	//Initiate Connection with Database

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Error")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT * FROM poloniex_historic_data where date=$1`
	err = db.QueryRow(sqlStatement).Scan(&date)

}

//Returns Poloniex Chart Data
//Parameters : Currency pair, Start time , End time

func (p *Poloniex) getChartData(currencyPair string, start string, end string) {

	//Initialise connection with database

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("Error")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//Get URL from constants.go

	url := poloniexBaseURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	//Append the user defined parameters to the url

	q := req.URL.Query()
	q.Add("command", "returnChartData")
	q.Add("currencyPair", currencyPair)
	q.Add("start", start)
	q.Add("end", end)
	req.URL.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	//Get the data from API and convert the data to byte array

	res, _ := p.client.Do(request)

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	//Store the data to charData struct

	var data chartData
	json.Unmarshal(body, &data)
	fmt.Printf("Results: %v\n", data)

	//Loop over the entire data and store it in the table
	for i := range data.Result {
		sqlStatement := `INSERT INTO poloniex_chart_data(date,high,low,open,close,volume,quoteVolume,weightedAverage) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
		_, err = db.Exec(sqlStatement, data.Result[i].Date, data.Result[i].High, data.Result[i].Low, data.Result[i].Open, data.Result[i].Close, data.Result[i].Volume, data.Result[i].QuoteVolume, data.Result[i].WeightedAverage)

	}

}
