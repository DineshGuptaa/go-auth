// PATH: go-auth/models/index.go

package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

var DB *gorm.DB

func InitDB(cfg Config) {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
          SlowThreshold:              time.Second,   // Slow SQL threshold
          LogLevel:                   logger.Info, // Log level
          IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
          ParameterizedQueries:      true,           // Don't include params in the SQL log
          Colorful:                  true,          // Disable color
        },
      )
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
    if err != nil {
        panic(err)
    }
    //db.AutoMigrate(false)
    // if err := db.AutoMigrate(&User{}); err != nil {
    //     panic(err)
    // }
    // log.Println("Migrated database")

    DB = db
}
