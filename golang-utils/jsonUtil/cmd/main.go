package main

import (
	"fmt"
	"io"
	"net/http"

	"com/jsonutil/internal/client"
	"com/jsonutil/internal/util"
)

func main() {
	fmt.Println("I'm running !!")
	//cpuMetricsLocation := os.GetEnv("CPU_JSON_FILE_LOCATION")
	cpuMetricsLocation := "C:\\Users\\vepeddada\\Downloads\\cupquota3.tf.json"

	//projectId := os.Getenv("PROJECT_ID")
	projectId := "totemic-chalice-345108"
	resp, err := client.FetchMetericsData(projectId)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(bodyBytes))

		output, err := util.Parse(bodyBytes)
		if err != nil {
			fmt.Println(err)
		}

		if output != nil {
			util.WriteCpuQuoteJson(output, projectId, cpuMetricsLocation)
		}
	}

}
