package iptodate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const CF_API_URL = "https://www.cloudflare.com/api_json.html"

func SetAddress(ip string, domain string, name string, email string, id string, token string) (bool, error) {
	v := url.Values{}
	v.Set("tkn", token)
	v.Set("email", email)
	v.Set("a", "rec_edit")
	v.Set("type", "A")
	v.Set("z", domain)
	v.Set("id", id)
	v.Set("name", name)
	v.Set("content", ip)
	v.Set("service_mode", "0")
	resp, err := http.PostForm(CF_API_URL, v)
	if err != nil {
		errorMessage := fmt.Sprintf("I was unable to post to CF: %s", err)
		return false, errors.New(errorMessage)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("I was unable to read CF Response: %s", err)
		return false, errors.New(errorMessage)
	}
	var response CFUpdateResponse
	json.Unmarshal(body, &response)
	if response.Result != "success" {
		errorMessage := fmt.Sprintf("I got an error from CF: %s", response.Message)
		return false, errors.New(errorMessage)
	}
	return true, nil
}

type CFUpdateResponse struct {
	Result  string `json:"result"` //This is all we need, really
	Message string `json:"msg"`
}

func GetId(domain string, email string, token string, name string) (string, error) {
	v := url.Values{}
	v.Set("tkn", token)
	v.Set("email", email)
	v.Set("z", domain)
	v.Set("a", "rec_load_all")
	resp, err := http.PostForm(CF_API_URL, v)
	if err != nil {
		errorMessage := fmt.Sprintf("I was unable to post to CF: %s", err)
		return "", errors.New(errorMessage)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("I was unable to read CF Response: %s", err)
		return "", errors.New(errorMessage)
	}
	var response CloudflareDomainLists
	json.Unmarshal(body, &response)
	if response.Result != "success" {
		errorMessage := fmt.Sprintf("I got an error from CF: %s", response.Message)
		return "", errors.New(errorMessage)
	}
	for _, v := range response.Response.Recs.Results {
		if strings.ToLower(v.DisplayName) == name {
			return v.RecID, nil
		}
	}
	return "", errors.New("I couldn't find that name, create a new one with the -a flag!")
}

type CloudflareDomainLists struct {
	Response struct {
		Recs struct {
			Count   float64 `json:"count"`
			HasMore bool    `json:"has_more"`
			Results []struct {
				RecID          string `json:"rec_id"`
				Content        string `json:"content"`
				DisplayContent string `json:"display_content"`
				DisplayName    string `json:"display_name"`
				Name           string `json:"name"`
				Type           string `json:"type"`
				ZoneName       string `json:"zone_name"`
			} `json:"objs"`
		} `json:"recs"`
	} `json:"response"`
	Result  string `json:"result"`
	Message string `json:"msg"`
}
