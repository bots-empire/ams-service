package service

import (
	"go.uber.org/zap"

	"ams-service/internal/db/accesses"
)

type Manager struct {
	storage accesses.Implementation

	whiteList []int64

	logger *zap.Logger
}

func NewManager(logger *zap.Logger, storage accesses.Implementation, whiteList []int64) *Manager {
	return &Manager{
		storage:   storage,
		whiteList: whiteList,
		logger:    logger,
	}
}
