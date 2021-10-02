package service

import "context"

// TodoService exposes all available use cases of Todo domain.
type TodoService interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
}

// CreateTodoRequest represents a request for creating a todo.
type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateTodoResponse represents a response of a request to create a todo.
type CreateTodoResponse struct{}
