package main

import (
	"github/quocdaitrn/todos/app/service/serviceimpl"
	"github/quocdaitrn/todos/app/transport/reqresp/http/httpimpl"
	"github/quocdaitrn/todos/infra/repoimpl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "admin:acb@4123@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tRepo := repoimpl.NewTodoRepo(db)
	tSvc := serviceimpl.NewTodoService(tRepo)

	httpimpl.SetupHTTPServer(tSvc)
}
