package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chinnaxs/go_beer_client/internal/pkg/beverage"
)

const defaultBaseUrl = "http://localhost:8080/beers"

type ApiClient struct {
	HttpClient *http.Client
	BaseUrl    string
}

func (a *ApiClient) do(r *http.Request) (*http.Response, error) {
	return a.HttpClient.Do(r)
}

func NewDefaultApiClient() *ApiClient {
	return &ApiClient{
		HttpClient: &http.Client{},
		BaseUrl:    defaultBaseUrl,
	}
}

func (a *ApiClient) ListBeers() ([]beverage.Beer, error) {
	req, err := http.NewRequest("GET", a.BaseUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned %d", a.BaseUrl, resp.StatusCode)
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

func (a *ApiClient) GetBeer(name string) (*beverage.Beer, error) {
	url := fmt.Sprintf("%s/%s", a.BaseUrl, name)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.do(req)
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

func (a *ApiClient) UpdateBeer(beer *beverage.Beer) error {
	putBody, err := json.Marshal(beer)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s", a.BaseUrl, beer.Name)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(putBody))
	if err != nil {
		return err
	}
	resp, err := a.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("PUT %s returned %d", a.BaseUrl, resp.StatusCode)
	}
	return nil
}

func (a *ApiClient) DeleteBeer(name string) error {
	url := fmt.Sprintf("%s/%s", a.BaseUrl, name)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := a.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("DELETE %s returned %d", url, resp.StatusCode)
	}
	return nil
}
