package repoimpl

import (
	"context"
	"github/quocdaitrn/todos/domain/entity"
	"github/quocdaitrn/todos/domain/repo"
	"gorm.io/gorm"
)

// TodoRepo implements all methods of Todo repo.
type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) repo.TodoRepo {
	return &TodoRepo{db: db}
}

// InsertOne
func (r *TodoRepo) InsertOne(_ context.Context, t *entity.Todo) (int64, error) {
	tx := r.db.Create(t)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return tx.RowsAffected, nil
}
