package main

//TODO: Вынести в отдельный микросервис миграций?

import (
	appService "chickchirick-migrator/cmd/service"
	"chickchirick-migrator/config/dto"
	"chickchirick-migrator/console/service"
	factory "chickchirick-migrator/factory"
	"fmt"
)

func main() {
	factory.InitViper()

	dbConfig := dto.NewConfiguration()
	dbDecorator := appService.InitORM(&dbConfig)
	defer dbDecorator.CloseDB()

	commandService := service.CommandService{
		MigratorService: service.MigratorService{
			DBDecorator: dbDecorator,
		},
	}

	err := commandService.ParseInput()
	if err != nil {
		panic(fmt.Errorf("failed to parse input: %w", err))
	}

}
