package service

import (
	"context"
	"github.com/bots-empire/ams-service/model"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/bots-empire/ams-service/internal/entity"
)

func (m *Manager) GetAdminsID(ctx context.Context, query *entity.AdminsQuery) ([]int64, error) {
	m.logger.Info("get admins", zap.Any("query", query))

	model.TotalGetAdmins.WithLabelValues(
		query.Code,
	).Inc()
	ids, err := m.storage.GetUsersByQuery(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "get from db")
	}

	return ids, nil
}

func (m *Manager) GetAllAdmins(ctx context.Context) ([]*entity.Access, error) {
	m.logger.Info("get all admins")

	ids, err := m.storage.GetAllAccess(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get from db")
	}

	return ids, nil
}
