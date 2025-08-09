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
	cfg.PostgreSQL.URL = os.Getenv("DATABASE_URL")

	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// New Database
	db, err := databases.NewPostgresqlDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	g.UseDB(db)

	g.ApplyBasic(entities.User{}, entities.Customer{}, entities.CustomerAddress{}, entities.ConfigConstant{})
	g.ApplyInterface(func(queriers.Querier) {}, entities.Customer{})
	// g.ApplyBasic(
	// 	g.GenerateModel("users"),
	// )
	g.Execute()

	s := servers.NewServer(cfg, db)
	s.Start()
}
