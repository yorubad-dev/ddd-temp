package persistent

import (
	"log"
	"os"

	"github.com/KingDaemonX/ddd-template/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

const (
	connectingDB   string = "🌀 Attempting To Connect Application Database"
	errDBConn      string = "🚨 Error Occur While Connecting To Database"
	successfulConn string = "😎 Database Connected SuccessFully"
	migrationErr   string = "❎ Error Occur While Migrating Database Schema"
)

type ProjectInfrast struct {
	ProjectName repository.ProjectRepository
	db          *gorm.DB
}

func ConnectDatabase() *ProjectInfrast {
	log.Println(connectingDB)

	dsn := os.Getenv("DATABASE_CONN_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(errDBConn)
		log.Fatalf("Error : %s", err.Error())
	}

	Conn = db
	if err := autoMigrate(Conn); err != nil {
		log.Println(migrationErr)
		log.Fatalf("Error : %s", err.Error())
	}

	return &ProjectInfrast{
		ProjectName: NewProjectInfra(db),
		db:          db,
	}
}
