package service

import (
	"ams-service/internal/entity"
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (m *Manager) AddIncomeInfo(ctx context.Context, add *entity.IncomeInfo) error {
	m.logger.Info("income info", zap.Any("income info", add))

	incInfo, err := m.storage.GetIncomeInfoByID(ctx, add.UserID, add.TypeBot)
	if err != nil {
		return errors.Wrap(err, "failed get income info from db")
	}

	err = m.storage.SaveIncomeInfo(ctx, incInfo)
	if err != nil {
		return errors.Wrap(err, "save income info in db")
	}

	return nil
}
