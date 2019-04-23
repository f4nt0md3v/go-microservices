// Package storage implement work with SQLite: connection/models/key-value storage/migration.
package cockroach

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-microservices/libs/config"
	"go-microservices/libs/logger"
	"os"
	"sync"
)

var database_connect *gorm.DB
var connMutexR = sync.RWMutex{}

// Return SQLite connection.
func Connection() (*gorm.DB, error) {
	connMutexR.Lock()
	defer connMutexR.Unlock()
	if database_connect == nil {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("postgresql://%s:%s@%s:%d?sslmode=%s&dbname=%s",
				config.GetString("cockroach_user"),
				config.GetString("cockroach_pass"),
				config.GetString("cockroach_host"),
				config.GetInt("cockroach_port"),
				config.GetString("cockroach_sslmode"),
				config.GetString("cockroach_db")))
		if err != nil {
			logger.GetCockroach().Error(10, map[string]interface{}{
				"error": err,
			})
			os.Exit(1)
			return nil, err
		}

		logger.GetCockroach().Info(10)
		database_connect = db
		migrate(db)
	}
	return database_connect, nil
}
