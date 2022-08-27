package accesses

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"ams-service/internal/entity"
)

const getAccessQuery = `SELECT user_id, code, additional, user_name, user_first_name, user_last_name
FROM ams.accesses
WHERE user_id = $1 AND code = $2;`

func (s *Storage) GetByCode(ctx context.Context, userID int64, code string) (*entity.Access, error) {
	rows, err := s.db.Query(
		ctx,
		getAccessQuery,
		userID,
		code,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get access by user_id")
	}
	defer rows.Close()

	acss, err := readAccessFromRows(rows)
	if err != nil {
		return nil, errors.Wrap(err, "scan from rows")
	}

	if len(acss) == 0 {
		return nil, nil
	}

	return acss[0], nil
}

func readAccessFromRows(rows pgx.Rows) ([]*entity.Access, error) {
	allAccess := make([]*entity.Access, 0)

	for rows.Next() {
		access := &entity.Access{}

		if err := rows.Scan(
			&access.UserID,
			&access.Code,
			pq.Array(&access.Additional),
			&access.UserName,
			&access.UserFirstName,
			&access.UserLastName,
		); err != nil {
			return nil, errors.Wrap(err, "scan access")
		}

		allAccess = append(allAccess, access)
	}

	return allAccess, nil
}
