package factory

import (
	mainService "chickchirick-migrator/cmd/service"
	"chickchirick-migrator/config"
	"chickchirick-migrator/migrator"
	"chickchirick-migrator/migrator/dto"
	migratorService "chickchirick-migrator/migrator/service"
	fakerService "chickchirick-migrator/pkg/chirik_faker/service"
	"github.com/spf13/viper"
)

type migratorOptions struct { //Конфигурация структуры
	createIndexAfterCreateTable bool
}

type MigratorOption func(options *migratorOptions)

func WithCreateIndexAfterCreateTable() MigratorOption { //Функция конфигурации,
	return func(mOptions *migratorOptions) {
		mOptions.createIndexAfterCreateTable = true
	}
}

func InitMigrator(dBDecorator mainService.DBDecorator, opts ...MigratorOption) migrator.Migrator {
	var mOptions migratorOptions
	for _, opt := range opts {
		opt(&mOptions)
	}

	mConfig := dto.MConfig{
		CreateIndexAfterCreateTable: mOptions.createIndexAfterCreateTable,
		MigrationFilesPath:          viper.GetString(config.MigrationPath),
		EnableTableNamespace:        viper.GetBool(config.EnableTableNamespace),
		FixtureCount:                viper.GetInt(config.FixtureCount),
		FixturePrefix:               viper.GetString(config.FixturePrefix),
		FixtureNilColumns:           fakerService.ParseExcludeString(viper.GetString(config.FixtureNilColumns)),
	}

	mDiContainer := dto.MigratorDiContainer{
		MConfig:     mConfig,
		DBDecorator: dBDecorator,
	}

	tCreator := migratorService.TableCreator{
		SqlProcessor: migratorService.SqlProcessor{MigratorDiContainer: mDiContainer},
	}
	fCreator := migratorService.FixtureCreator{
		SqlProcessor: migratorService.SqlProcessor{MigratorDiContainer: mDiContainer},
	}

	return migrator.Migrator{
		TableCreator:   tCreator,
		FixtureCreator: fCreator,
	}
}
