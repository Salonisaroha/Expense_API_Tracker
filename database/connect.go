package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

type DB struct {
	endpoint string
	name     string
}

func NewDBInstance(endpoint, name string) *DB {
	return &DB{
		endpoint: endpoint,
		name:     name,
	}
}

func (d *DB) Connect() {
	if d.endpoint == "" || d.name == "" {
		panic("invalid database credentials")
	}

	dsn := d.buildConnectionString()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	Connection = db
	fmt.Println("Connected To Database ")
}

func (d *DB) buildConnectionString() string {
	return fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", d.endpoint, d.name)
}
