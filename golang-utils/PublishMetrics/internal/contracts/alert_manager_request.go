package contracts

import "time"

type AlertManagerRequest struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alert           `json:"alerts"`
	CommonLabels      CommonLabels      `json:"commonLabels"`
	CommonAnnotations CommonAnnotations `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
}

type Alert struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     time.Time   `json:"startsAt"`
	EndsAt       time.Time   `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
	Fingerprint  string      `json:"fingerprint"`
}

type Labels struct {
	Alertname string `json:"alertname"`
	Env       string `json:"env"`
	Instance  string `json:"instance"`
	Severity  string `json:"severity"`
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type CommonLabels struct {
	Alertname string `json:"alertname"`
	Env       string `json:"env"`
	Instance  string `json:"instance"`
	Severity  string `json:"severity"`
}

type CommonAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
