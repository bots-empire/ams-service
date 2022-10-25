package accesses

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/bots-empire/ams-service/internal/entity"
)

const updateOrCreateAccessQuery = `INSERT INTO ams.accesses (user_id, code, additional, user_name, user_first_name, user_last_name) 
VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (user_id, code) DO UPDATE SET additional = $3, user_name = $4, user_first_name = $5, user_last_name = $6;`

func (s *Storage) Update(ctx context.Context, access *entity.Access) error {
	result, err := s.db.Exec(
		ctx,
		updateOrCreateAccessQuery,
		access.UserID,
		access.Code,
		pq.Array(access.Additional),
		access.UserName,
		access.UserFirstName,
		access.UserLastName,
	)
	if err != nil {
		return errors.Wrap(err, "failed update or create")
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
