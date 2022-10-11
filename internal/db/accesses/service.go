package accesses

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"ams-service/internal/entity"
)

type Implementation interface {
	GetAll(ctx context.Context, userID int64) ([]*entity.Access, error)
	GetByCode(ctx context.Context, userID int64, code string) (*entity.Access, error)
	Update(ctx context.Context, access *entity.Access) error

	GetUsersByQuery(ctx context.Context, access *entity.AdminsQuery) ([]int64, error)

	GetAllAccess(ctx context.Context) ([]*entity.Access, error)
}

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(connect *pgxpool.Pool) *Storage {
	return &Storage{
		db: connect,
	}
}
