package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "http://localhost:8080/beers"

func ListBeers(c *http.Client) ([]Beer, error) {
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
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

func GetBeer(c *http.Client, name string) (*Beer, error) {
	url := fmt.Sprintf("%s/%s", baseUrl, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
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

func UpdateBeer(c *http.Client, beer *Beer) error {
	putBody, err := json.Marshal(beer)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", baseUrl, bytes.NewReader(putBody))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return fmt.Errorf("PUT %s returned %d", baseUrl, resp.StatusCode)
	}
	return nil
}

func DeleteBeer(c *http.Client, name string) error {
	url := fmt.Sprintf("%s/%s", baseUrl, name)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return fmt.Errorf("DELETE %s returned %d", url, resp.StatusCode)
	}
	return nil
}
