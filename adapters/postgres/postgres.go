package postgres

import (
	"context"
	"database/sql"
	_ "embed"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func NewGORMDB(postgresConfig Config, logger *zap.Logger) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(postgresConfig.Addr))
}

func NewSQLDB(orm *gorm.DB) (*sql.DB, error) {
	return orm.DB()
}

func HookConnection(lifecycle fx.Lifecycle, db *sql.DB, logger *zap.Logger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return db.PingContext(ctx) },
		OnStop:  func(ctx context.Context) error { return db.Close() },
	})
}

func CreateTables(db *sql.DB) error {
	_, err := db.ExecContext(context.Background(), schema)
	return err
}
