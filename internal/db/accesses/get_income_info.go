package accesses

import (
	"ams-service/internal/entity"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const getIncomeInfoQuery = `SELECT user_id, bot_link, bot_name, income_source, type_bot
FROM ams.income_info WHERE user_id = $1 AND type_bot = $2;`

func (s *Storage) GetIncomeInfoByID(ctx context.Context, userID int64, botType string) ([]*entity.IncomeInfo, error) {
	rows, err := s.db.Query(
		ctx,
		getIncomeInfoQuery,
		userID,
		botType,
	)
	if err != nil {
		return nil, errors.Wrap(err, "get income info by user_id")
	}

	defer rows.Close()

	incomeInfo, err := readIncomeInfoRows(rows)
	if err != nil {
		return nil, errors.Wrap(err, "failed read income info")
	}

	return incomeInfo, nil
}

func readIncomeInfoRows(rows pgx.Rows) ([]*entity.IncomeInfo, error) {
	allIncomeInfo := make([]*entity.IncomeInfo, 0)

	for rows.Next() {
		incomeInfo := &entity.IncomeInfo{}

		if err := rows.Scan(
			&incomeInfo.UserID,
			&incomeInfo.BotLink,
			&incomeInfo.BotName,
			&incomeInfo.IncomeSource,
			&incomeInfo.TypeBot,
		); err != nil {
			return nil, errors.Wrap(err, "scan rows income info")
		}

		allIncomeInfo = append(allIncomeInfo, incomeInfo)
	}

	return allIncomeInfo, nil
}
