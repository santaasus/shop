package dbservice

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	DataBase DataBaseConfig `json:"Database"`
}

type DataBaseConfig struct {
	Postgres PostgresConfig `json:"Postgres"`
}

type PostgresConfig struct {
	ShopGo ShopGoConfig `json:"ShopGo"`
}

type ShopGoConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	User     string `json:"user"`
	DbName   string `json:"db_name"`
	Password string `json:"password"`
	SSLMode  string `json:"sslmode"`
}

func Connect() (*sql.DB, error) {
	// var config Config

	// data, err := os.ReadFile("config.json")
	// if err != nil {
	// 	// log.Fatalf("Error reading json: %v", err)
	// 	return nil, err
	// }

	// err = json.Unmarshal(data, &config)
	// if err != nil {
	// 	// log.Fatalf("Error unmarshal json: %v", err)
	// 	return nil, err
	// }

	// username := "user=" + config.DataBase.Postgres.ShopGo.User
	// dbName := "dbname=" + config.DataBase.Postgres.ShopGo.DbName
	// dbPassword := "password=" + config.DataBase.Postgres.ShopGo.Password
	// host := "host=" + config.DataBase.Postgres.ShopGo.Host
	// port := "port=" + strconv.Itoa(config.DataBase.Postgres.ShopGo.Port)
	// sslMode := "sslmode=" + config.DataBase.Postgres.ShopGo.SSLMode

	// dbConfig := username + " " + dbName + " " + dbPassword + " " + host + " " + port + " " + sslMode

	dataBaseUrl := "user=vladimirslesarev password=mypassword dbname=mydb host=localhost port=5432 sslmode=disable" //os.Getenv("DATABASE_URL")
	if dataBaseUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not assign")
	}

	db, err := sql.Open("postgres", dataBaseUrl)
	if err != nil {
		// log.Fatalf("Error sql open: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		// log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	return db, nil
}
