package accesses

import (
	"ams-service/internal/entity"
	"context"
	"github.com/pkg/errors"
)

const GetAllIncomeInfoQuery = `SELECT user_id, bot_link, bot_name, income_source, type_bot
FROM ams.income_info;`

func (s *Storage) GetAllIncomeInfo(ctx context.Context) ([]*entity.IncomeInfo, error) {
	rows, err := s.db.Query(
		ctx,
		GetAllIncomeInfoQuery,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bad request to get all income info")
	}
	defer rows.Close()

	return readIncomeInfoRows(rows)
}
