package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	queries "github.com/jolorunyomi-convoso/auto-audit/sql/queries"
)

type config struct {
	MySQL struct {
		User         string
		Password     string
		DatabaseName string
		Host         string
		Port         int
	}
}

func loadConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}
	cfg.MySQL.User = os.Getenv("MYSQL_USER")
	cfg.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	cfg.MySQL.DatabaseName = os.Getenv("MYSQL_DBNAME")
	cfg.MySQL.Host = os.Getenv("MYSQL_HOST")
	portStr := os.Getenv("MYSQL_PORT")
	cfg.MySQL.Port, err = strconv.Atoi(portStr)
	if err != nil {
		return cfg, err
	}
	if cfg.MySQL.Host == "" && cfg.MySQL.Port == 0 {
		return cfg, fmt.Errorf("no MYSQL config provided")
	}

	return cfg, nil
}

func run(cfg config) error {
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=false", cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port, cfg.MySQL.DatabaseName)
	db, err := sql.Open("mysql", connection_string)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		return err
	}

	var (
		pa_name   string
		pa_type   string
		ch_entity string
		ch_name   string
		ch_type   string
	)

	retrieveChildItems := queries.RetrieveChildItems()
	stmt, err := db.Prepare(retrieveChildItems)
	if err != nil {
		return err
	}
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pa_name, &pa_type, &ch_entity, &ch_name, &ch_type)
		if err != nil {
			return err
		}
		fmt.Printf("Retrieved data from db: %s % s%s %s %v \n", pa_name, pa_type, ch_entity, ch_name, ch_type)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	err = rows.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	err = run(cfg)
	if err != nil {
		panic(err)
	}

}
