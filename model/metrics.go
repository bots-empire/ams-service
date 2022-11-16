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
)
