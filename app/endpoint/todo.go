package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github/quocdaitrn/todos/app/service"
)

type TodoServiceEndpoints struct {
	CreateTodoEndpoint endpoint.Endpoint
}

func NewTodoServiceEndpoints(svc service.TodoService) *TodoServiceEndpoints {
	epts := &TodoServiceEndpoints{}
	epts.CreateTodoEndpoint = newCreateTodoEndpoint(svc)

	return epts
}

func newCreateTodoEndpoint(svc service.TodoService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return svc.CreateTodo(ctx, req.(*service.CreateTodoRequest))
	}
}
