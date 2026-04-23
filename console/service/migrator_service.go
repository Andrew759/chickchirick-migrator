package service

import (
	"chickchirick-migrator/cmd/service"
	"chickchirick-migrator/migrator/factory"
	migratorDto "chickchirick-migrator/migrator/provider"
)

type MigratorService struct {
	DBDecorator service.DBDecorator
}

func (ms MigratorService) DoMigrate(migratorEntities map[string][]migratorDto.MigratorInfo) error {
	migrator := factory.InitMigrator(ms.DBDecorator, factory.WithCreateIndexAfterCreateTable())
	return migrator.Crea(migratorEntities)
}
