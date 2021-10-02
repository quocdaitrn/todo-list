package serviceimpl

import (
	"context"

	"github/quocdaitrn/todos/app/service"
	"github/quocdaitrn/todos/domain/entity"
	"github/quocdaitrn/todos/domain/repo"
)

// TodoService implements TodoService.
type TodoService struct {
	tRepo repo.TodoRepo
}

func NewTodoService(tRepo repo.TodoRepo) service.TodoService {
	return &TodoService{
		tRepo: tRepo,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *service.CreateTodoRequest) (*service.CreateTodoResponse, error) {
	todo := &entity.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	_, err := s.tRepo.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &service.CreateTodoResponse{}, nil
}
