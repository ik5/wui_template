package metrics

import (
	"context"
	"time"
)

// MetricDetails a list of metrics details
type MetricDetails []string

// MetricDetail holds information about a metric
type MetricDetail struct {
	Ctx       context.Context `json:"context,omitempty"`
	Name      string          `json:"name"`
	Times     uint64          `json:"times"`
	Details   MetricDetails   `json:"details,omitempty"`
	StartedAt time.Time       `json:"started_at"`
	EndedAt   time.Time       `json:"ended_at"`
}

// Metrics is a map for each type of action
type Metrics map[string]MetricDetail

var currentMetrics Metrics

const emptyTime = "0001-01-01 00:00:00 +0000 UTC"

// SetMetricsCtx set new content to the current metric including context
func SetMetricsCtx(ctx *context.Context, name string, details MetricDetails) {
	if currentMetrics == nil {
		currentMetrics = make(Metrics)
	}

	metric, found := currentMetrics[name]

	// sets the latest context
	if ctx != nil {
		metric.Ctx = *ctx
	}
	if found {
		metric.EndedAt = time.Now()
		if len(metric.Details) == 0 {
			metric.Details = append(metric.Details, details...)
		} else {
			metric.Details = details
		}
		metric.Times++
	} else {
		metric.Name = name
		metric.Details = details
		metric.Times = 1
		metric.StartedAt = time.Now()
		metric.EndedAt = time.Now()
	}

	currentMetrics[name] = metric

}

// SetMetrics set new content to current metric
func SetMetrics(name string, details MetricDetails) {
	SetMetricsCtx(nil, name, details)
}
