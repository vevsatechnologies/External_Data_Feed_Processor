package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://api.decred.org/?c=gsd"
)

type POS struct {
	client *http.Client
}

type POSData struct {
	APIEnabled           string  `json:"APIEnabled"`
	APIVersionsSupported string  `json:"APIVersionsSupported"`
	Network              string  `json:"Network"`
	Url                  string  `json:"URL"`
	Launched             float64 `json:"Launched"`
	LastUpdated          float64 `json:"LastUpdated"`
	Immature             float64 `json:"Immature"`
	Live                 string  `json: "Live"`
	Voted                string  `json:"Voted"`
	Missed               string  `json:"Missed"`
	PoolFees             string  `json:"PoolFees"`
	ProportionLive       float64 `json:"ProportionLive"`
	ProportionMissed     float64 `json: "ProportionMissed"`
	UserCount            int     `json:"UserCount"`
	UserCountActive      int     `json:"UserCountActive"`
}

type POSID struct {
	Ethan    map[string]*POSData `json:"Ethan"`
	Papa     map[string]*POSData `json:"Papa"`
	Echo     map[string]*POSData `json:"Echo"`
	Mike     map[string]*POSData `json:"Mike"`
	Oscar    map[string]*POSData `json:"Oscar"`
	Golf     map[string]*POSData `json:"Golf"`
	November map[string]*POSData `json:"November"`
	Ray      map[string]*POSData `json:"Ray"`
	Hotel    map[string]*POSData `json:"Hotel"`
	Juliett  map[string]*POSData `json:"Juliett"`
	Lima     map[string]*POSData `json:"Lima"`
	India    map[string]*POSData `json:"India"`
	Bravo    map[string]*POSData `json:"Bravo"`
	Delta    map[string]*POSData `json:"Delta"`
	Kilo     map[string]*POSData `json:"Kilo"`
}

func (p *POS) getPOS() {
	url := url
	request, err := http.NewRequest("GET", url, nil)

	res, _ := p.client.Do(request)

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data POSID
	json.Unmarshal(body, &data)

	fmt.Printf("Results: %v\n", data)

	//Loop over the entire list to insert data into the table
	for i := 0; i < 15; i++ {
		sqlStatement := `INSERT INTO POSDataTable(POSId,APIEnabled,APIVersionsSupported,Network,Url,Launched,LastUpdated,Immature,Live,Voted,Missed,PoolFees,ProportionLive,ProportionMissed,UserCount,UserCountActive,timestamp) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`
		_, err = db.Exec(sqlStatement, data)

	}
}
