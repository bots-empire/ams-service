package service

import (
	"context"
	"github.com/bots-empire/ams-service/internal/model"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/bots-empire/ams-service/internal/entity"
)

func (m *Manager) CheckAccess(ctx context.Context, check *entity.Access) (bool, error) {
	m.logger.Info("check access", zap.Any("access", check))

	model.TotalCheckedAccesses.WithLabelValues(
		strconv.FormatInt(check.UserID, 10),
		check.UserFirstName,
	).Inc()

	if m.checkWhiteList(check.UserID) {
		return true, nil
	}

	acs, err := m.storage.GetByCode(ctx, check.UserID, check.Code)
	if err != nil {
		return false, errors.Wrap(err, "get from db")
	}

	if acs == nil {
		return false, nil
	}

	if len(acs.Additional) == 0 {
		return true, nil
	}

	if len(check.Additional) == 0 {
		return false, nil
	}

	for _, rule := range check.Additional {
		if !contains(acs.Additional, rule) {
			return false, nil
		}
	}

	return true, nil
}

func (m *Manager) checkWhiteList(userID int64) bool {
	for _, uID := range m.whiteList {
		if uID == userID {
			return true
		}
	}

	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (m *Manager) AddAccess(ctx context.Context, add *entity.Access) error {
	m.logger.Info("add access", zap.Any("access", add))

	model.TotalAddedAccesses.WithLabelValues(
		strconv.FormatInt(add.UserID, 10),
		add.UserFirstName,
	).Inc()

	acs, err := m.storage.GetByCode(ctx, add.UserID, add.Code)
	if err != nil {
		return errors.Wrap(err, "get from db")
	}

	if acs != nil {
		acs.MergeAccess(add)
	} else {
		acs = add
	}

	err = m.storage.Update(ctx, acs)
	if err != nil {
		return errors.Wrap(err, "update access in db")
	}

	return nil
}

func (m *Manager) DepriveAccess(ctx context.Context, check *entity.Access) error {
	// TODO: implement method
	return nil
}
