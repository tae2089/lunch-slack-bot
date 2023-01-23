package config

import (
	"bc-labs-lunch-bot/ent"
	"bc-labs-lunch-bot/ent/migrate"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"sync"
	"time"
)

var (
	onceEnt sync.Once
	client  *ent.Client
)

type DbConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
	Database string `env:"DB_NAME,required"`
}

func OpenDB() *ent.Client {
	log.Println("start OpenDB")
	if client == nil {
		onceEnt.Do(func() {
			log.Println("init OpenDB")
			db, err := sql.Open("pgx", GetDbUrl())
			if err != nil {
				log.Println(err)
				client.Close()
				return
			}
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(100)
			db.SetConnMaxLifetime(time.Minute)
			// Create an ent.Driver from `db`.
			drv := entsql.OpenDB(dialect.Postgres, db)
			client = ent.NewClient(ent.Driver(drv))
			err = client.Schema.Create(context.Background(),
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
			)
			if err != nil {
				log.Println("error - ", err)
			}
		})
	}
	return client
}
