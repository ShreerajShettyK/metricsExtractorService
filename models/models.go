package models

type SplunkMetric struct {
	Time       int64                  `json:"time,omitempty"`
	Event      string                 `json:"event"`
	Host       string                 `json:"host"`
	Source     string                 `json:"source"`
	Sourcetype string                 `json:"sourcetype"`
	Fields     map[string]interface{} `json:"fields"`
	Index      string                 `json:"index"`
}
