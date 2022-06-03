package conf

import (
	"flag"
	"os"
)

type Mongo struct {
	Username string
	Password string
	Host     string
	Database string
}

var (
	Database Mongo
	Port     *string
)

func Init() {
	Port = flag.String("p", "3000", "Port")
	flag.Parse()

	Database = Mongo{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MOGNO_HOST"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
}
