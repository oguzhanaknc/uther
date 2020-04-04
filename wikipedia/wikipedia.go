package wikipedia

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/buger/jsonparser"
)

//GetWiki : parametre olarak gelen değeri wikipedia üzerinde arayıp sonucu döndürüyor
func GetWiki(search string) (string, bool) {

	url := "https://tr.wikipedia.org/w/api.php?action=query&prop=extracts&exsentences=5&exlimit=1&titles=" + search + "&explaintext=1&formatversion=2&format=json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var results []map[string]interface{}
	data := []byte(body)
	result, _, _, err := jsonparser.Get(data, "query", "pages")
	if err != nil {
		return "Üzgünüm internette &&&& ile ilgili bir sonuç bulamadım", false
	}
	errr := json.Unmarshal([]byte(result), &results)
	if errr != nil {
		return "Üzgünüm internette &&&& ile ilgili bir sonuç bulamadım", false
	}
	if results[0]["extract"] != nil {
		return results[0]["extract"].(string), true
	}
	return "Üzgünüm internette &&&& ile ilgili bir sonuç bulamadım", false
}
