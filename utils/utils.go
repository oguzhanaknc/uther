package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
)

type hava struct {
	max         string
	min         string
	icon        string
	day         string
	description string
}

//Havadurumu : ankara havadurumu bilgisini döndürüyor
func Havadurumu() (string, string, string, string, string) {
	Hava := new(hava)
	url := "https://api.collectapi.com/weather/getWeather?data.lang=tr&data.city=ankara"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "apikey 73AH5ed5RqZdJ54FN3Y92h:0tWV1lYRv0990s9Anm0DLK")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// Declared an empty interface of type Array
	var results []map[string]interface{}
	data := []byte(body)
	result, _, _, _ := jsonparser.Get(data, "result")
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(result), &results)
	Hava.day = results[0]["day"].(string)
	Hava.icon = results[0]["icon"].(string)
	Hava.max = results[0]["max"].(string)
	Hava.min = results[0]["min"].(string)
	Hava.description = results[0]["description"].(string)
	return Hava.day, Hava.description, Hava.icon, Hava.max, Hava.min
}
