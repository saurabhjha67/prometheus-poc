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
	token := "ya29.A0ARrdaM8ZuONgnjSzqb6GLiFR4hExAgxdLfyLrjuMnxrf19BwmBikkybPvSVQO1gpT5hupgEndFZ1Tf7rGS6mghzBBcO1hIxnKnDr2ucCyHdiElCDmsp2dXrLObEtWM7cCTKoA5f4Oow7FKJzWXtwkRoRWCS1IdL0OAw4QA"
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
