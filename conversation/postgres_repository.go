package conversation

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func AutoMigrate(lifecycle fx.Lifecycle, db *gorm.DB) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return db.WithContext(ctx).AutoMigrate(Conversation{})
		},
	})
}
