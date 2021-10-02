package repo

import (
	"context"
	"github/quocdaitrn/todos/domain/entity"
)

// TodoRepo provides all methods to interact with Todo entity.
type TodoRepo interface {
	InsertOne(ctx context.Context, t *entity.Todo) (int64, error)
}
