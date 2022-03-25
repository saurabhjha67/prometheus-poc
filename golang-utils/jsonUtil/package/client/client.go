package client

import (
	"fmt"
	"net/http"
	"strings"
)

type Client interface {
	CallMetericsInfo() (*http.Response, error)
}

type client struct {
	apiUrl     string
	httpClient *http.Client
}

func NewClient(projectId string) (Client, error) {
	//apiUrl := os.Getenv("API_URL")
	apiUrl := "https://serviceusage.googleapis.com/v1beta1/projects/<projectId>/services/compute.googleapis.com/consumerQuotaMetrics"
	apiUrl = strings.Replace(apiUrl, "<projectId>", projectId, 1)

	return &client{apiUrl: apiUrl,
		httpClient: &http.Client{}}, nil
}

func (c *client) CallMetericsInfo() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.apiUrl, nil)

	if err != nil {
		return nil, err
	}

	//token := os.Getenv("TOKEN")
	token := "ya29.A0ARrdaM_765V2GAIYHrd-QDaeP8PjUNZLCMjjNPY_YYYzWz0Hc2QdP4D5c6ywJJqvqwYoUozt70LkAxJ_wZkZ1QvgKCKjZMl-x7rEBMZjDR2oslIPY7VHe5A4iy3VHZLuv402I2BtGw7R_a1zNOdRFKmZLboToroFRWTl-g"
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
