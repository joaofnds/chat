package conversation

import (
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}
