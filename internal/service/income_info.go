package service

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/bots-empire/ams-service/internal/entity"
)

func (m *Manager) AddIncomeInfo(ctx context.Context, add *entity.IncomeInfo) error {
	m.logger.Info("income info", zap.Any("income info", add))

	err := m.storage.SaveIncomeInfo(ctx, add)
	if err != nil {
		return errors.Wrap(err, "save income info in db")
	}

	return nil
}

func (m *Manager) GetIncomeInfo(ctx context.Context, id int64, typeBot string) (*entity.IncomeInfo, error) {
	m.logger.Info("income info", zap.Any("income info", id))

	incInfo, err := m.storage.GetIncomeInfoByID(ctx, id, typeBot)
	if err != nil {
		return nil, errors.Wrap(err, "get income info in db")
	}

	return incInfo, nil
}
