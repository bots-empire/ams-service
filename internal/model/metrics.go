package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalAddedIncome = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_added_income_users",
			Help: "Total count of added income users",
		},
		[]string{"bot_link", "bot_name"},
	)

	TotalGetIncome = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_get_income_users",
			Help: "Total count of get income users",
		},
		[]string{"user_id", "bot_type"},
	)

	TotalCheckedAccesses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_checked_accesses",
			Help: "Total count of checked accesses",
		},
		[]string{"user_id", "user_first_name"},
	)

	TotalAddedAccesses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_added_accesses",
			Help: "Total count of added accesses",
		},
		[]string{"user_id", "user_first_name"},
	)

	TotalGetAdmins = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_get_admins",
			Help: "Total count of get admins",
		},
		[]string{"query_code"},
	)

	ResponseTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "response_time",
			Help:    "Response time of route",
			Buckets: []float64{0.1, 0.5, 0.9, 1, 5, 10},
		},
		[]string{"handler"},
	)
)
