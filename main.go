package main

import (
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	fetchHistoricData(1, "") //parameters : exchangeID,currency pair, start time, end time
	for {
		getHistoricData(1, "BTC-DCR", "1514764800", "1514851200")   //parameters : exchangeID,currency pair, start time, end time
		getTickData("BTC_DCR")                                      //parameters :  Currency pair
		getPoloniexChartData("BTC_DCR", "1405699200", "9999999999") //parameters: Currency Pair, start time , end time
	}

}

// Exchange id = 0 for Poloneix
// Exchange id =1 for Bittrex

func fetchHistoricData(exchangeID int, date string) {

	if exchangeID == 0 { //Poloniex exchange
		user := Poloniex{

			client: &http.Client{},
		}
		user.fetchPoloniexData(date)

	}

	if exchangeID == 1 { //Bittrex Exchange

		user := Bittrex{
			client: &http.Client{},
		}
		user.fetchBittrexData(date)
	}

}

// Exchange id = 0 for Poloneix
// Exchange id =1 for Bittrex

func getHistoricData(exchangeID int, currencyPair string, startTime string, endTime string) {

	if exchangeID == 0 { //Poloniex exchange
		user := Poloniex{

			client: &http.Client{},
		}
		user.getPoloniexData(currencyPair, startTime, endTime)

	}

	if exchangeID == 1 { //Bittrex Exchange

		user := Bittrex{
			client: &http.Client{},
		}
		user.getBittrexData(currencyPair)
	}

	//Time delay of 24 hours

	time.Sleep(86400 * time.Second)
}

//Get ticker data from Bittrex exchange

func getTickData(currencyPair string) {

	user := Bittrex{
		client: &http.Client{},
	}
	user.getTicks(currencyPair)

}

//Get chart data from Poloniex Exchange
func getPoloniexChartData(currencyPair string, startTime string, endTime string) {
	user := Poloniex{

		client: &http.Client{},
	}
	user.getChartData(currencyPair, startTime, endTime)

}
