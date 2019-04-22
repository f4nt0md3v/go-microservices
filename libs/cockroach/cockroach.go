// Package storage implement work with SQLite: connection/models/key-value storage/migration.
package cockroach

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-microservices/libs/config"
	"go-microservices/libs/logger"
)

var database_connect *gorm.DB

// Return SQLite connection.
func Connection() (*gorm.DB, interface{}) {
	if database_connect == nil {
		db, err := gorm.Open("postgres",
			fmt.Sprintf("postgresql://%s:%s@%s:%d?sslmode=%s",
				config.GetString("cockroach_user"),
				config.GetString("cockroach_password"),
				config.GetString("cockroach_host"),
				config.GetInt("cockroach_port"),
				config.GetString("cockroach_sslmode")))
		if err != nil {
			if err != nil {
				logger.GetCockroach().Error(10, map[string]interface{}{
					"error": err,
				})
				return nil, err
			}
			return nil, err
		}

		logger.GetCockroach().Info(10)
		database_connect = db
		migrate(db)
	}
	return database_connect, nil
}
