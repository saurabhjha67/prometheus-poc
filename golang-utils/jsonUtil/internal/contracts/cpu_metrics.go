package contracts

type CPUMetrics struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Metric      string `json:"metric"`
	Unit        string `json:"unit"`
}

type Metrics struct {
	MetricsAll []CPUMetrics `json:"metrics"`
}
