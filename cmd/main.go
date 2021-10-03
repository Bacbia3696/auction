package main

import (
	"database/sql"
	"log"

	db "github.com/bacbia3696/auction/db/sqlc"
	"github.com/bacbia3696/auction/internal/config"
	"github.com/bacbia3696/auction/internal/constant"
	"github.com/bacbia3696/auction/internal/server"
	"github.com/bacbia3696/auction/internal/util"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := loadConfig()
	configLogger(cfg)
	conn := loadDB(cfg)
	s := server.New(cfg, db.NewStore(conn))
	log.Fatal(s.Serve())
}

func configLogger(cfg *config.Config) {
	if cfg.Environment == constant.PRODUCTION {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetLevel(logrus.DebugLevel)
}

func loadConfig() *config.Config {
	cfg, err := config.Load(".")
	if err != nil {
		panic(err)
	}
	if cfg.Environment != constant.PRODUCTION {
		util.DebugPrint("config", cfg)
	}
	return cfg
}

func loadDB(cfg *config.Config) *sql.DB {
	conn, err := sql.Open("postgres", cfg.DbSource)
	if err != nil {
		panic(err)
	}
	return conn
}
