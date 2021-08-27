package repo

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	sql "github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-quiz-api/internal/models"
)

var ErrCannotAddEntity = errors.New("cannot add entity")
var NotExists = errors.New("entity does not exist")

// Repo - интерфейс хранилища для сущности Quiz
type Repo interface {
	AddEntity(ctx context.Context, entity models.Quiz) (uint64, error)
	AddEntities(ctx context.Context, entities []models.Quiz) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Quiz, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Quiz, error)
	RemoveEntity(ctx context.Context, entityId uint64) bool
}

type dbRepo struct {
	stmBuilder  sq.StatementBuilderType
	tableName   string
	userId      string
	classroomId string
	link        string
}

func (d *dbRepo) AddEntity(ctx context.Context, entity models.Quiz) (uint64, error) {
	query := d.stmBuilder.Insert(d.tableName).
		Columns(d.userId, d.classroomId, d.link).
		Suffix("RETURNING id").
		Values(entity.UserId, entity.ClassroomId, entity.Link)
	newQuizId := uint64(0)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return newQuizId, err
	}
	rows.Next()
	if err := rows.Scan(&newQuizId); err != nil {
		return newQuizId, err
	}

	return newQuizId, nil
}

func (d *dbRepo) AddEntities(ctx context.Context, entities []models.Quiz) error {
	query := d.stmBuilder.Insert(d.tableName).
		Columns(d.userId, d.classroomId, d.link)

	for _, r := range entities {
		query = query.Values(r.UserId, r.ClassroomId, r.Link)
	}

	if _, err := query.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}

func (d *dbRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Quiz, error) {
	query := d.stmBuilder.Select(fmt.Sprintf("id, %v, %v, %v", d.userId, d.classroomId, d.link)).
		From(d.tableName).
		Offset(offset). //not the fastest approach but will keep as is in favor of simplicity (ability to remove objects makes it a bit complex)
		Limit(limit)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	requests := make([]models.Quiz, 0, limit)
	for rows.Next() {
		req := models.Quiz{}
		if err := rows.Scan(&req.Id, &req.UserId, &req.ClassroomId, &req.Link); err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func (d *dbRepo) DescribeEntity(ctx context.Context, entityId uint64) (*models.Quiz, error) {
	query := d.stmBuilder.Select("id", d.userId, d.classroomId, d.link).
		From(d.tableName).
		Where("id=?", entityId)

	row, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	entity := models.Quiz{}
	if !row.Next() {
		return nil, NotExists
	}
	if err := row.Scan(&entity.Id, &entity.UserId, &entity.ClassroomId, &entity.Link); err != nil {
		return nil, err
	}
	return &entity, nil
}

func (d *dbRepo) RemoveEntity(ctx context.Context, entityId uint64) bool {
	query := d.stmBuilder.Delete(d.tableName).Where("id=?", entityId)

	rows, err := query.ExecContext(ctx)
	if err != nil {
		return false
	}

	rowsDeleted, err := rows.RowsAffected()
	if err != nil {
		return false
	}
	return rowsDeleted > 0

}

func NewRepo(db *sql.DB) *dbRepo {
	stmCache := sq.NewStmtCache(db)

	return &dbRepo{
		stmBuilder:  sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(stmCache),
		tableName:   "quiz",
		userId:      "user_id",
		classroomId: "classroom_id",
		link:        "link",
	}

}
