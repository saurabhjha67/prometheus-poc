package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"com/jsonutil/internal/contracts"
)

func Parse(metricJsonContent []byte) ([]string, error) {
	data := contracts.Metrics{}
	output := []string{}

	_ = json.Unmarshal(metricJsonContent, &data)

	for i := 0; i < len(data.MetricsAll); i++ {
		if strings.Contains(strings.ToUpper(data.MetricsAll[i].DisplayName), "CPU") && !strings.Contains(strings.ToUpper(data.MetricsAll[i].DisplayName), "COMMITTED") && !strings.Contains(strings.ToUpper(data.MetricsAll[i].DisplayName), "PREEMPTIBLE") && data.MetricsAll[i].DisplayName != "N2D CPUs" {
			output = append(output, data.MetricsAll[i].Metric)
		}
	}
	fmt.Println(output)
	return output, nil
}

func WriteCpuQuoteJson(output []string, projectId string, cpuMetricsLocation string) {
	cpu := contracts.CPU{
		CPUNames:     output,
		Override:     "4",
		OverrideDisk: "1000",
		ProjectId:    projectId,
		Region:       "asia-east1",
		Source:       "../cpuquota",
	}

	module := contracts.Module{
		CPUs: []contracts.CPU{
			cpu,
		},
	}

	cpuQuota := contracts.CPUQuota{
		Modules: []contracts.Module{
			module,
		},
	}
	outputBytes, _ := json.MarshalIndent(cpuQuota, "", " ")
	fmt.Println(string(outputBytes))
	err := os.WriteFile(cpuMetricsLocation, outputBytes, 0644)
	if err != nil {
		fmt.Println("error persisting into file ", err)
	}
}
