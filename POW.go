package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vevsatechnologies/External_Data_Feed_Processor/models"
)

const (
	url1 = "http://api.f2pool.com/decred/address"
	url2 = "https://www2.coinmine.pl/dcr/index.php?page=api&action=getpoolstatus"
	url3 = "http://mining.luxor.tech/API/DCR/stats/"
	url4 = "https://dcr.suprnova.cc/index.php?page=api&action=getpoolstatus"
	url5 = "https://altpool.net/api/status"
	url6 = "https://altpool.net/api/currencies"
	url7 = "https://pool.mn/dcr/index.php?page=api&action=getpoolstatus"
	url8 = "https://dcred.eu/api/api?q=network"
)

type POW struct {
	client *http.Client
}

type POWdata struct {
	date          string              `json : "date"`
	hashper       string              `json : "hashper" `
	blocksper     string              `json:"blocksper"`
	luck          string              `json:"luck"`
	miners        string              `json:"miners"`
	pphash        string              `json:"pphash"`
	ppshare       string              `json:"ppshare"`
	totalKickback string              `json:"total_kickback"`
	price         float64             `json:"price"`
	hashrate      float64             `json:"hashrate"`
	blocksfound   int64               `json:"blocksfound"`
	totalminers   int64               `json:"totalminers"`
	globalStats   []globalStatsValues `json:"globalStats"`
	dataVal       dataVal             `json:"data"`
	decred        altpool             `json:"decred"`
	dcr           altpoolCurrency     `json:"DCR"`
	success       string              `json:"success"`
	lastUpdate    int64               `json:"lastUpdate"`
	mainnet       mainnet             `json:"mainnet"`
	blockReward   blockReward         `json:"blockReward"`
}

type mainnet struct {
	currentHeight     int64 `json:"currentHeight"`
	networkHashrate   int64 `json:"networkHashrate"`
	networkDifficulty int64 `json:"networkDifficulty"`
}

type blockReward struct {
	total float64 `json:"total"`
	pow   float64 `json:"pow"`
	pos   float64 `json:"pos"`
	dev   float64 `json:"dev"`
}

type globalStatsValues struct {
	time              string  `json:"time"`
	networkHashrate   float64 `json:"network_hashrate"`
	poolHashrate      float64 `json:"pool_hashrate"`
	workers           int64   `json:"workers"`
	networkDifficulty float64 `json:"network_difficulty"`
	coinPrice         float64 `json:"coin_price"`
	btcPrice          float64 `json:"btc_price"`
}

type dataVal struct {
	poolName            string  `json:"pool_name"`
	hashrate            float64 `json:"hashrate"`
	efficiency          float64 `json:'efficiency"`
	progress            float64 `json:"progress"`
	workers             int64   `json:"workers"`
	currentnetworkblock int64   `json:"currentnetworkblock"`
	nextnetworkblock    int64   `json:"nextnetworkblock"`
	lastblock           int64   `json:"lastblock"`
	networkdiff         float64 `json:"networkdiff"`
	esttime             float64 `json:"esttime"`
	estshares           int64   `json:"estshares"`
	timesincelast       int64   `json:"timesincelast"`
	nethashrate         int64   `json:"nethashrate"`
}

type altpool struct {
	name             string  `json:"name"`
	port             int64   `json:"port"`
	coins            int64   `json:"coins"`
	fees             int64   `json:"fees"`
	hashrate         int64   `json:"hashrate"`
	workers          int64   `json:"workers"`
	estimate_current float64 `json:"estimate_current"`
	estimate_last24h float64 `json"estimate_last24h"`
	actual_last24h   float64 `json:"actual_last24h"`
	mbtc_mh_factor   int64   `json:"mbtc_mh_factor"`
	hashrate_last24h float64 `json:"hashrate_last24h"`
	rental_current   float64 `json:"rental_current"`
}

type altpoolCurrency struct {
	algo          string  `json:"algo"`
	port          int64   `json:"port"`
	name          string  `json:"name"`
	height        int64   `json:"height"`
	workers       int64   `json:"workers"`
	shares        int64   `json:"shares"`
	hashrate      int64   `json:"hashrate"`
	estimate      float64 `json:"estimate"`
	blocks24h     int64   `json:"24h_blocks"`
	btc24h        float64 `json:"24h_btc"`
	lastblock     int64   `json:"lastblock"`
	timesincelast int64   `json:"timesincelast"`
}

func (p *POW) getPOW(id int, url string, api_key string) {

	req, err := http.NewRequest("GET", url, nil)

	if len(api_key) != 0 {
		q := req.URL.Query()
		q.Add("api_key", api_key)
		req.URL.RawQuery = q.Encode()
	}

	request, err := http.NewRequest("GET", req.URL.String(), nil)

	res, _ := p.client.Do(request)

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var data POWdata
	json.Unmarshal(body, &data)

	fmt.Println(string(body))

	fmt.Printf("Results: %v\n", data)

	//Loop over the entire list to insert data into the table
	for i := 0; i < 15; i++ {

		var p1 models.Powdatatable

		p1.ID = id
		p1.HashRate = data.hashrate | data.dataVal.hashrate | data.pphash
		p1.Efficiency = data.dataVal.efficiency
		p1.Progress = data.dataVal.progress
		p1.Workers = data.globalStats[0].workers
		p1.CurrentNetworkBlock = data.dataVal.currentnetworkblock
		p1.NextNetworkBlock = data.dataVal.nextnetworkblock
		p1.LastBlock = data.dataVal.lastblock
		p1.NetworkDiff = data.dataVal.networkdiff | data.globalStats[0].networkDifficulty | data.mainnet.networkDifficulty
		p1.EstTime = data.globalStats[0].time | data.dataVal.esttime
		p1.EstShare = data.dataVal.estshares
		p1.TimeSinceLast = data.dataVal.timesincelast
		p1.NetHashrate = data.globalStats[0].networkHashrate
		p1.BlocksFound = data.blocksfound
		p1.TotalMiners = data.totalminers
		p1.Time = data.globalStats[0].time
		p1.NetworkDifficulty = data.globalStats[0].networkDifficulty
		p1.CoinPrice = data.globalStats[0].coinPrice
		p1.BTCPrice = data.globalStats[0].btcPrice
		p1.Est = data.dcr.estimate
		p1.Date = data.date
		p1.Blocksper = data.blocksper
		p1.Luck = data.luck
		p1.Ppshare = data.ppshare
		p1.Totalkickback = data.totalKickback
		p1.Success = data.success
		p1.Lastupdate = data.lastUpdate
		p1.Name = data.decred.name
		p1.Port = data.decred.port
		p1.Fees = data.decred.fees
		p1.EstimateCurrent = data.decred.estimate_current
		p1.EstimateLast24h = data.decred.estimate_last24h
		p1.Actual24H = data.decred.actual_last24h
		p1.MBTCMHFactor = data.decred.mbtc_mh_factor
		p1.HashrateLast24h = data.decred.hashrate_last24h
		p1.RentalCurrent = data.decred.rental_current
		p1.Height = data.dcr.height
		p1.Blocks24h = data.dcr.blocks24h
		p1.BTC24H = data.dcr.btc24h
		p1.Currentheight = data.mainnet.currentHeight
		p1.Total = data.blockReward.total
		p1.Pos = data.blockReward.pos
		p1.Pow = data.blockReward.pow
		p1.Dev = data.blockReward.dev
		p1.Powid = id

		err := p1.Insert(db)
	}

}
