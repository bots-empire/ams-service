package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	//Income
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

	//Accesses

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
	//Users

	TotalGetAdmins = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_get_admins",
			Help: "Total count of get admins",
		},
		[]string{"query_code"},
	)

	//response time

	TimeUserRoutes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "time_user_routes",
			Help:    "Response time of user route",
			Buckets: []float64{0.1, 0.5, 0.9, 1},
		},
		[]string{"handler"},
	)

	TimeAccessesRoutes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "time_accesses_routes",
			Help:    "Response time of accesses route",
			Buckets: []float64{0.1, 0.5, 0.9, 1},
		},
		[]string{"handler"},
	)

	TimeIncomeRoutes = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "time_income_info_routes",
			Help:    "Response time of income info route",
			Buckets: []float64{0.1, 0.5, 0.9, 1},
		},
		[]string{"handler"},
	)
)
