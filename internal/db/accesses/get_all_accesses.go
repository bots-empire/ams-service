package accesses

import (
	"context"

	"github.com/pkg/errors"

	"ams-service/internal/entity"
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
