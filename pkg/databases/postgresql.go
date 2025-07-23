package databases

import (
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/sweetcandy273/go-teerawut/configs"
	"github.com/sweetcandy273/go-teerawut/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresqlDBConnection new postgresql database connection
func NewPostgresqlDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(postgresUrl), &gorm.Config{})
	if err != nil {
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return db, nil
}
