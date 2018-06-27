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
	POSID map[string]*POSData `json:"Ethan"`
	// POSID map[string]*POSData `json:"Papa"`
	// POSID map[string]*POSData `json:"Echo"`
	// POSID map[string]*POSData `json:"Mike"`
	// POSID map[string]*POSData `json:"Oscar"`
	// POSID map[string]*POSData `json:"Golf"`
	// POSID map[string]*POSData `json:"November"`
	// POSID map[string]*POSData `json:"Ray"`
	// POSID map[string]*POSData `json:"Hotel"`
	// POSID map[string]*POSData `json:"Juliett"`
	// POSID map[string]*POSData `json:"Lima"`
	// POSID map[string]*POSData `json:"India"`
	// POSID map[string]*POSData `json:"Bravo"`
	// POSID map[string]*POSData `json:"Delta"`
	// POSID map[string]*POSData `json:"Kilo"`
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
	for i := range data.POSID {
		sqlStatement := `INSERT INTO POSData(POSId,APIEnabled,APIVersionsSupported,Network,Url,Launched,LastUpdated,Immature,Live,Voted,Missed,PoolFees,ProportionLive,ProportionMissed,UserCount,UserCountActive,timestamp) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`
		_, err = db.Exec(sqlStatement, data.POSID[0], data.POSID[1], data.POSID[2], data.POSID[3], data.POSID[4], data.POSID[5], data.POSID[6], data.POSID[7], data.POSID[8], data.POSID[9], data.POSID[10], data.POSID[11], data.POSID[12], data.POSID[13], data.POSID[14], Date(NOW()))

	}
}
