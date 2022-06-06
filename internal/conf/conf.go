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
	LOKI     *bool
	LOKI_URI string
	Port     *string
)

func Init() {
	LOKI_URI = os.Getenv("CN_LOKI_URI")

	Database = Mongo{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
		Host:     os.Getenv("MONGO_HOST"),
		Database: os.Getenv("MONGO_DATABASE"),
	}

	Port = flag.String("p", "3000", "Port")
	LOKI = flag.Bool("l", true, "Nutze Loki")
	flag.Parse()
}
