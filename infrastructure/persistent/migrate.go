package persistent

import (
	"github.com/KingDaemonX/ddd-template/domain/entity"
	"gorm.io/gorm"
)

func autoMigrate(conn *gorm.DB) error {
	return conn.AutoMigrate(&entity.Entity{})
}
