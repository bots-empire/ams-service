package accesses

import (
	"context"

	"github.com/pkg/errors"

	"ams-service/internal/entity"
)

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
