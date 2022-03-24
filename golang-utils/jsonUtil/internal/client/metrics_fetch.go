package client

import (
	"com/jsonutil/package/client"
	"fmt"
	"net/http"
)

func FetchMetericsData(projectId string) (*http.Response, error) {
	httpClient, err := client.NewClient(projectId)

	if err != nil {
		fmt.Println("error", err)
	}

	response, error := httpClient.CallMetericsInfo()

	if error != nil {
		fmt.Println("error", err)
	}

	return response, error
}
