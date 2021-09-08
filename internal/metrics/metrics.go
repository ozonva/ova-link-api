package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	CreateSuccessResponseCounter()
	MultiCreateSuccessResponseCounter()
	UpdateSuccessResponseCounter()
	ListSuccessResponseCounter()
	DescribeSuccessResponseCounter()
	RemoveSuccessResponseCounter()
}

type metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	updateCounter      prometheus.Counter
	listCounter        prometheus.Counter
	removeCounter      prometheus.Counter
	describeCounter    prometheus.Counter
}

func (m *metrics) CreateSuccessResponseCounter() {
	m.createCounter.Inc()
}

func (m *metrics) MultiCreateSuccessResponseCounter() {
	m.multiCreateCounter.Inc()
}

func (m *metrics) UpdateSuccessResponseCounter() {
	m.updateCounter.Inc()
}

func (m *metrics) ListSuccessResponseCounter() {
	m.listCounter.Inc()
}

func (m *metrics) RemoveSuccessResponseCounter() {
	m.removeCounter.Inc()
}

func (m *metrics) DescribeSuccessResponseCounter() {
	m.describeCounter.Inc()
}

func NewMetrics() Metrics {
	m := &metrics{
		createCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_create_response",
			Help: "Api create success response",
		}),
		updateCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_update_response",
			Help: "Api update success response",
		}),
		listCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_list_response",
			Help: "Api list success response",
		}),
		removeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_remove_response",
			Help: "Api remove success response",
		}),
		multiCreateCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_multi_create_response",
			Help: "Api remove success response",
		}),
		describeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "success_describe_response",
			Help: "Api describe success response",
		}),
	}

	prometheus.MustRegister(m.createCounter)
	prometheus.MustRegister(m.updateCounter)
	prometheus.MustRegister(m.listCounter)
	prometheus.MustRegister(m.removeCounter)
	prometheus.MustRegister(m.multiCreateCounter)
	prometheus.MustRegister(m.describeCounter)

	return m
}
