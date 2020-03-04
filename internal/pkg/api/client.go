package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chinnaxs/go_beer_client/internal/pkg/beverage"
)

const baseUrl = "http://localhost:8080/beers"

func ListBeers(c *http.Client) ([]beverage.Beer, error) {
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
	var beers []beverage.Beer
	err = json.Unmarshal(getBody, &beers)
	if err != nil {
		return nil, err
	}
	return beers, nil
}

func GetBeer(c *http.Client, name string) (*beverage.Beer, error) {
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned %d", url, resp.StatusCode)
	}
	getBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var beer beverage.Beer
	err = json.Unmarshal([]byte(getBody), &beer)
	if err != nil {
		return nil, err
	}
	return &beer, nil
}

func UpdateBeer(c *http.Client, beer *beverage.Beer) error {
	putBody, err := json.Marshal(beer)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s", baseUrl, beer.Name)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(putBody))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
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
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("DELETE %s returned %d", url, resp.StatusCode)
	}
	return nil
}
