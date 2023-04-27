package application

import "github.com/KingDaemonX/ddd-template/domain/repository"

type ProjectApp struct {
	Project repository.ProjectRepository
}

type ProjectApplication interface {
	// implement thesame thing as the repository because it serves as a passage between the interface and domain
}

// create passage method here
var _ repository.ProjectRepository = &ProjectApp{}
