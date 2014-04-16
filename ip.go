package iptodate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var response Address
	json.Unmarshal(body, &response)
	return response.Ip, nil
}

type Address struct {
	Ip string `json:"ip"`
}
