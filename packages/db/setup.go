package db

import (
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Import file driver
	_ "github.com/lib/pq"                                // Import PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// setup database connection
func SetupDatabase() {
	DB = initGormConnection()
	// if DB == nil {
	// 	return
	// }
	// DB.AutoMigrate(types.User{})
	// DB.AutoMigrate(types.Permission{})
	// DB.AutoMigrate(types.Invitation{})
	// DB.AutoMigrate(types.Transaction{})
	// DB.AutoMigrate(types.AccessKey{})
	// DB.AutoMigrate(types.Thread{})
	// DB.AutoMigrate(types.ThreadMember{})
	// DB.AutoMigrate(types.Message{})
}

// generate connection string from  environment variables
func generateConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	return connectionStr
}

func initGormConnection() *gorm.DB {
	connectionString := generateConnectionString()
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("database connected")
		return db
	}

}
