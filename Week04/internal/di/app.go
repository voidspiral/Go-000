package di

import (
	service "school/internal/service"
)

type App struct {
	grpc *service.School
}

func NewApp(grpc *service.School) *App {
	return &App{grpc: grpc}
}

func (app *App) Start() error {
	return app.grpc.Start()
}

func (app *App) Stop() {
	app.grpc.Stop()
}
