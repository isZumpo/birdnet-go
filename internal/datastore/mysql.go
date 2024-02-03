package datastore

import (
	"fmt"
	"log"

	"github.com/tphakala/birdnet-go/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLStore implements DataStore for MySQL
type MySQLStore struct {
	DataStore
	Ctx *conf.Context
}

func validateMySQLConfig(ctx *conf.Context) error {
	// Add validation logic for MySQL configuration
	// Return an error if the configuration is invalid
	return nil
}

// InitializeDatabase sets up the MySQL database connection
func (store *MySQLStore) Open() error {
	if err := validateMySQLConfig(store.Ctx); err != nil {
		return err // validateMySQLConfig returns a properly formatted error
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		store.Ctx.Settings.Output.MySQL.Username, store.Ctx.Settings.Output.MySQL.Password,
		store.Ctx.Settings.Output.MySQL.Host, store.Ctx.Settings.Output.MySQL.Port,
		store.Ctx.Settings.Output.MySQL.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open MySQL database: %v\n", err)
		return fmt.Errorf("failed to open MySQL database: %v", err)
	}

	store.DB = db
	return performAutoMigration(db, store.Ctx.Settings.Debug, "MySQL", dsn)
}

// SaveToDatabase inserts a new Note record into the SQLite database
func (store *MySQLStore) Save(note Note) error {
	if store.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	if err := store.DB.Create(&note).Error; err != nil {
		log.Printf("Failed to save note: %v\n", err)
		return err
	}

	return nil
}

// Close MySQL database connections
func (store *MySQLStore) Close() error {
	// Ensure that the store's DB field is not nil to avoid a panic
	if store.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// Retrieve the generic database object from the GORM DB object
	sqlDB, err := store.DB.DB()
	if err != nil {
		log.Printf("Failed to retrieve generic DB object: %v\n", err)
		return err
	}

	// Close the generic database object, which closes the underlying SQL database connection
	if err := sqlDB.Close(); err != nil {
		log.Printf("Failed to close MySQL database: %v\n", err)
		return err
	}

	return nil
}