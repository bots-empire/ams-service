package accesses

import (
	"ams-service/internal/entity"
	"context"
	"github.com/pkg/errors"
)

const saveIncomeInfoQuery = `INSERT INTO ams.income_info (user_id, bot_link, bot_name, income_source, type_bot) 
VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (user_id, bot_link) DO UPDATE SET bot_name = $3, income_source = $4, type_bot = $5;`

func (s *Storage) SaveIncomeInfo(ctx context.Context, info *entity.IncomeInfo) error {
	_, err := s.db.Exec(
		ctx,
		saveIncomeInfoQuery,
		info.UserID,
		info.BotLink,
		info.BotName,
		info.IncomeSource,
		info.TypeBot,
	)
	if err != nil {
		return errors.Wrap(err, "bad request to save income info")
	}

	return nil
}
