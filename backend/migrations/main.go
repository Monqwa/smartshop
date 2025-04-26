package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

var dir = flag.String("dir", "./migrations", "directory with migration files")

func main() {
	flag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt("db.port"),
		viper.GetString("db.sslmode"),
	)

	// Now "pgx" is a known driver
	db, err := goose.OpenDBWithDriver("pgx", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}
	defer db.Close()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	cmd, extra := args[0], args[1:]
	if err := goose.RunContext(context.Background(), cmd, db, *dir, extra...); err != nil {
		log.Fatalf("goose %s: %v", cmd, err)
	}
}
