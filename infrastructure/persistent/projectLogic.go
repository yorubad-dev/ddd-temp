package persistent

import (
	"github.com/KingDaemonX/ddd-template/domain/repository"
	"gorm.io/gorm"
)

type ProjectInfra struct {
	database *gorm.DB
}

func NewProjectInfra(db *gorm.DB) ProjectInfra {
	return ProjectInfra{database: db}
}

var _ repository.ProjectRepository = &ProjectInfra{}
