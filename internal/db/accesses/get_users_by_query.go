package accesses

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"ams-service/internal/entity"
)

const getUserIDsQuery = `SELECT user_id, additional
FROM ams.accesses
WHERE code = $1;`

func (s *Storage) GetUsersByQuery(ctx context.Context, query *entity.AdminsQuery) ([]int64, error) {
	rows, err := s.db.Query(
		ctx,
		getUserIDsQuery,
		query.Code,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get access by user_id")
	}
	defer rows.Close()

	admins, err := readIDsFromRows(rows)
	if err != nil {
		return nil, errors.Wrap(err, "read ids from rows")
	}

	return getMatchIDs(admins, query.Additional), nil
}

func readIDsFromRows(rows pgx.Rows) ([]*entity.Access, error) {
	allAccess := make([]*entity.Access, 0)

	for rows.Next() {
		access := &entity.Access{}

		if err := rows.Scan(
			&access.UserID,
			pq.Array(&access.Additional),
		); err != nil {
			return nil, errors.Wrap(err, "scan access")
		}

		allAccess = append(allAccess, access)
	}

	return allAccess, nil
}

func getMatchIDs(admins []*entity.Access, check []string) []int64 {
	ids := make([]int64, 0)

	for _, admin := range admins {
		aditionals := make(map[string]struct{}, len(admin.Additional))

		for _, key := range admin.Additional {
			aditionals[key] = struct{}{}
		}

		match := true
		for _, role := range check {
			if _, exist := aditionals[role]; !exist {
				match = false
			}
		}

		if !match {
			continue
		}

		ids = append(ids, admin.UserID)
	}

	return ids
}
