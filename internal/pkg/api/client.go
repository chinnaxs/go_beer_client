package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "http://localhost:8080/beers/"

func ListAllBeers() ([]Beer, error) {
	resp, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned %d", baseUrl, resp.StatusCode)
	}
	getBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var beers []Beer
	err = json.Unmarshal([]byte(getBody), &beers)
	if err != nil {
		return nil, err
	}
	return beers, nil
}

func GetBeer(name string) (*Beer, error) {
	url := baseUrl + name
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned %d", url, resp.StatusCode)
	}
	getBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var beer Beer
	err = json.Unmarshal([]byte(getBody), &beer)
	if err != nil {
		return nil, err
	}
	return &beer, nil
}

func UpdateBeer(beer *Beer) error {
	postBody, err := json.Marshal(beer)
	if err != nil {
		return err
	}
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("POST %s returned %d", baseUrl, resp.StatusCode)
	}
	return nil
}

// func DeleteBeer(name string) error {
// 	url := baseUrl + name
// 	req, err := http.NewRequest("DELETE", url, nil)
// 	if err != nil {
// 		return err
// 	}
// 	resp, err := http.(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("DELETE %s returned %d", url, resp.StatusCode)
// 	}
// }
