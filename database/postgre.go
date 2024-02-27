package database

import (
	"fmt"
	"goskeleton/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(cfg *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, 
		
	}), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}


func DBMigration(db *gorm.DB) {
	//db.AutoMigrate(&model.Users{})	
}