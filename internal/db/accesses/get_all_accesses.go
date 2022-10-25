package accesses

import (
	"context"

	"github.com/pkg/errors"

	"github.com/bots-empire/ams-service/internal/entity"
)

func (s *Storage) GetAll(ctx context.Context, userID int64) ([]*entity.Access, error) {
	rows, err := s.db.Query(
		ctx,
		getAccessQuery,
		userID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get access by user_id")
	}
	defer rows.Close()

	return readAccessFromRows(rows)
}

const getAllAccessQuery = `SELECT user_id, code, additional, user_name, user_first_name, user_last_name
FROM ams.accesses;`

func (s *Storage) GetAllAccess(ctx context.Context) ([]*entity.Access, error) {
	rows, err := s.db.Query(
		ctx,
		getAllAccessQuery,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get all access")
	}
	defer rows.Close()

	return readAccessFromRows(rows)
}
