package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sweetcandy273/go-teerawut/configs"
	"github.com/sweetcandy273/go-teerawut/modules/entities"

	"github.com/sweetcandy273/go-teerawut/modules/servers"
	databases "github.com/sweetcandy273/go-teerawut/pkg/databases"
	"github.com/sweetcandy273/go-teerawut/pkg/queriers"
	"gorm.io/gen"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load("./.env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	g.UseDB(db)

	g.ApplyBasic(entities.User{}, entities.Customer{})
	g.ApplyInterface(func(queriers.Querier) {}, entities.Customer{})
	// g.ApplyBasic(
	// 	g.GenerateModel("users"),
	// )
	g.Execute()

	s := servers.NewServer(cfg, db)
	s.Start()
}
