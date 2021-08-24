package repo

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	sql "github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

var ErrCannotAddEntity = errors.New("cannot add entity")

// Repo - интерфейс хранилища для сущности Quiz
type Repo interface {
	AddEntities(ctx context.Context, entities []models.Quiz) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Quiz, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Quiz, error)
	RemoveEntity(ctx context.Context, entityId uint64) bool
}

type dbRepo struct {
	stmBuilder sq.StatementBuilderType
}

func (d *dbRepo) AddEntities(ctx context.Context, entities []models.Quiz) error {
	panic("implement me")
}

func (d *dbRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Quiz, error) {
	panic("implement me")
}

func (d *dbRepo) DescribeEntity(ctx context.Context, entityId uint64) (*models.Quiz, error) {
	panic("implement me")
}

func (d *dbRepo) RemoveEntity(ctx context.Context, entityId uint64) bool {
	panic("implement me")
}

func newRepo(db *sql.DB) *dbRepo {
	stmCache := sq.NewStmtCache(db)
	builder := sq.StatementBuilder
	builder.PlaceholderFormat(sq.Dollar).RunWith(stmCache)

	return &dbRepo{
		stmBuilder: builder,
	}

}
