package httpimpl

import (
	"github.com/go-chi/chi/v5"
	"github/quocdaitrn/todos/app/endpoint"
	"github/quocdaitrn/todos/app/service"
	"github/quocdaitrn/todos/app/transport/reqresp/http/codec"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func SetupHTTPServer(
	tSvc service.TodoService,

) {
	tSvcEpts := endpoint.NewTodoServiceEndpoints(tSvc)

	createTodoHandler := httptransport.NewServer(
		tSvcEpts.CreateTodoEndpoint,
		codec.DecodeCreateTodoRequest,
		codec.EncodeResponse,
	)

	r := chi.NewRouter()
	r.Method("POST", "/v1/rpc/create-todo", createTodoHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
