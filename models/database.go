package models

import (
	"fmt"

	"github.com/vishal2457/redmine-server/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the he database connection.
var DB *gorm.DB

// SetupDatabase migrates and sets up the database.
func SetupDatabase() {
	u := helper.GetEnv("DATABASE_USER", "root")
	p := helper.GetEnv("DATABSE_PASSWORD", "")
	h := helper.GetEnv("DATABASE_HOST", "localhost:3306")
	n := helper.GetEnv("DATABASE_NAME", "redmine")
	q := "charset=utf8mb4&parseTime=True&loc=Local"

	// Assemble the connection string.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", u, p, h, n, q)
	fmt.Print(dsn)
	// Connect to the database.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(
		&User{},
		&Project{},
		&ProjectUser{},
		&Task{},
	)

	if err != nil {
		panic("PANIC::cannot open database connection")
	} else {
		fmt.Print("INFO::Db connected")
	}

	DB = db
}
