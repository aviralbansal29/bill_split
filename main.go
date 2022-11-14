package main

import (
	"github.com/aviralbansal29/bill_split/config"
	"github.com/aviralbansal29/bill_split/config/database"
	_ "github.com/aviralbansal29/bill_split/docs/swagger"
	server "github.com/aviralbansal29/bill_split/transport"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

// @title Splitwise clone Service
// @version 1.0
// @BasePath /
func main() {
	config.InitiateGlobalInstance()
	database.Migrate()

	s := server.RestServer{}
	s.Setup()
	s.Start()
}
