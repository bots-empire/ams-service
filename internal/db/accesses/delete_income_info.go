package accesses

import (
	"context"
	"github.com/pkg/errors"
)

const DeleteIncomeInfoQuery = `DELETE FROM ams.income_info WHERE user_id = $1;`

func (s *Storage) DeleteIncomeInfo(ctx context.Context, userId int64) error {
	_, err := s.db.Exec(
		ctx,
		DeleteIncomeInfoQuery,
		userId,
	)
	if err != nil {
		return errors.Wrap(err, "delete income info")
	}

	return nil
}
