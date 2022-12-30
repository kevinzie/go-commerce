package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"kevinzie/go-commerce/pkg/utils"
	"log"
	"os"
	"strings"
	"time"
)

var Database *gorm.DB

func Connect() error {
	var err error
	// Build PostgreSQL connection URL.
	postgresConnURL, errorCon := utils.ConnectionURLBuilder("postgres")
	if errorCon != nil {
		return nil
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgresConnURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                            // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                           // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("ID", "id"), // use name replacer to change struct/field name before convert it to db name
		},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
	})

	if err != nil {
		return err
	}

	return nil
}
