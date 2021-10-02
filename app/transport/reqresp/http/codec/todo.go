package codec

import (
	"context"
	"encoding/json"
	"github/quocdaitrn/todos/app/service"
	"net/http"
)

func DecodeCreateTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	request := &service.CreateTodoRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}