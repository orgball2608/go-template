package mysql

import (
	"elearning/config"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math"
	"time"
)

type Database struct {
	*gorm.DB
}

func ConnectMysql(cfg *config.Environment) (*Database, error) {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		cfg.MysqlUserName, 
		cfg.MysqlPassword, 
		cfg.MysqlHost, 
		cfg.MysqlPort,
		cfg.MysqlDatabase,
	)
	for i := 0; i < 3; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger(cfg),
			// QueryFields:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			fmt.Printf("attempt connecting the database...(%d)\n", i+1)
			// Retry connecting DB
			time.Sleep(time.Second * time.Duration(math.Pow(3, float64(i))))
			continue
		}
		break
	}
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error dsn: %q", dsn))
	}
	return &Database{db}, nil
}

func (d Database) AutoMigrate() error {
	if err := addUserTable(d.DB); err != nil {
		return err
	}
	return nil

	//TODO: implement other tables
}
