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
	ID       *string
	Database Mongo
)

func Init() {
	ID = flag.String("s", "", "Computer id")
	flag.Parse()

	Database = Mongo{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MOGNO_HOST"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
}
