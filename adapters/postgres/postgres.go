package postgres

import (
	"context"
	"database/sql"
	_ "embed"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed schema.sql
var schema string

var Module = fx.Module(
	"postgres",
	fx.Provide(NewGORMDB),
	fx.Provide(NewSQLDB),
	fx.Provide(NewHealthChecker),
	fx.Invoke(HookConnection),
	fx.Invoke(CreateTables),
)

func NewGORMDB(postgresConfig Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(postgresConfig.Addr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
}

func NewSQLDB(orm *gorm.DB) (*sql.DB, error) {
	return orm.DB()
}

func HookConnection(lifecycle fx.Lifecycle, db *sql.DB) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return db.PingContext(ctx) },
		OnStop:  func(ctx context.Context) error { return db.Close() },
	})
}

func CreateTables(db *sql.DB) error {
	_, err := db.ExecContext(context.Background(), schema)
	return err
}
