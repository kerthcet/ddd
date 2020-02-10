package config

import (
	"os"
	"strconv"

	// load .env
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	ENV.init()
}

type dotEnv struct {
	ServiceName string
	DatabaseURL string
	KafkaHost   string
	GormLogMode bool
	GinMode     string
}

// ENV 所有环境变量
var ENV dotEnv

func (e *dotEnv) init() {
	e.DatabaseURL = os.Getenv("DATABASE_URL")
	if e.DatabaseURL == "" {
		panic("DATABASE_URL not exists")
	}

	e.KafkaHost = os.Getenv("KAFKA_HOST")
	if e.KafkaHost == "" {
		panic("KAFKA_HOST not exists")
	}

	logMode, err := strconv.ParseBool(os.Getenv("GORM_LOG_MODE"))
	if err != nil {
		panic("GORM_LOG_MODE is error")
	}
	e.GormLogMode = logMode

	e.GinMode = os.Getenv("GIN_MODE")
	if e.GinMode == "" {
		panic("GIN_MODE not exists")
	}
}
