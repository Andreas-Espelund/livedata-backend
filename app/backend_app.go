package app

import (
	"fmt"
	"github.com/Andreas-Espelund/livedata-backend/app/config"
	"github.com/Andreas-Espelund/livedata-backend/app/db"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

const (
	contextTimeout = 10 * time.Second
)

type Api struct {
	Logger *logrus.Entry
	Db     *sqlx.DB
	Config config.Config
}

func isLocal() bool {
	fmt.Print(os.Getenv("ENV"))
	return strings.ToLower(os.Getenv("ENV")) == "local"
}

func New() Api {

	conCfg := config.MSSQLConnectionConfig{
		SQLDialTimeout:               2,
		SQLConnectionTimeout:         2,
		SQLMaxOpenConnections:        50,
		SQLMaxIdleConnections:        15,
		SQLMaxConnectionLifetimeSecs: 300,
	}

	cfg := config.MSSQLDBConfig{
		SQLHost:               "livedata-db-server.database.windows.net",
		SQLDatabase:           "livedata-db",
		SQLPassword:           "SuperSecret1337",
		SQLUsername:           "sa",
		LocalSQLHost:          "localhost",
		LocalSQLPort:          "1436",
		LocalSQLDbName:        "livedata-db",
		MSSQLConnectionConfig: conCfg,
	}

	conf := config.Config{
		DBConfig: cfg,
		IsLocal:  isLocal,
	}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})

	logEntry := log.WithFields(logrus.Fields{
		"app": "livedata-backend",
		"env": "test",
	})

	database, err := db.NewDb(&conf)

	if err != nil {
		log.WithError(err)
	}
	return Api{
		Logger: logEntry,
		Db:     database,
		Config: conf,
	}
}
