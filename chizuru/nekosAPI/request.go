package nekosAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoRequest(target string) (res *Response, err error) {
	reqURL := fmt.Sprintf("%s/%s", baseUrl, target)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(Response)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func nekostext(target string) (*TEXT, error) {
	resp, err := http.Get(fmt.Sprintf("https://nekos.life/api/v2/%v", target))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var N = new(Response)
	json.NewDecoder(resp.Body).Decode(&N)
	return &N.Data.TEXT, nil
