package interfaces

import (
	"github.com/KingDaemonX/ddd-template/application"
)

type ProjectInterface struct {
	NameYouWant application.ProjectApplication
}

func NewProjectName(nyw application.ProjectApplication) ProjectInterface {
	return ProjectInterface{NameYouWant: nyw}
}

// write serializer method, pass to database and interact with the main file
