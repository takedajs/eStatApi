package estatapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../config"
)

type Estat struct {
	GET_STATS_DATA GetStatsData
}

type GetStatsData struct {
	RESULT           interface{}
	PARAMETER        interface{}
	STATISTICAL_DATA StatisticalData
}

type StatisticalData struct {
	RESULT_INF interface{}
	TABLE_INF  interface{}
	CLASS_INF  interface{}
	DATA_INF   DataInf
}

type DataInf struct {
	NOTE  interface{}
	VALUE []Value
}

type Value struct {
	Cat01  string `json:"@cat01"`
	Cat02  string `json:"@cat02"`
	Cat03  string `json:"@cat03"`
	Area   string `json:"@area"`
	Time   string `json:"@time"`
	Unit   string `json:"@unit"`
	Dollar string `json:"$"`
}

func Get() Estat {

	url := "http://api.e-stat.go.jp/rest/2.0/app/json/getStatsData?appId=" + config.AppId + "&statsDataId=0003104181"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var e Estat
	if err := json.Unmarshal([]byte(body), &e); err != nil {
		panic(err)
	}

	return e

}
