package persistence

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBEntry struct {
	Key        string `gorm:"primaryKey"`
	Value      string
	Expiration int64
}

var db *gorm.DB

// InitDB initializes the database connection.
func InitDB(dsn string, dbType string) error {
	var err error
	if dbType == "sqlite" {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else if dbType == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		return err
	}

	if err != nil {
		return err
	}

	// Migrate the schema
	return db.AutoMigrate(&DBEntry{})
}

// SaveToDB persists the current state of the store to the database.
func SaveToDB(data map[string]interface{}) error {
	if db == nil {
		return nil
	}

	for key, value := range data {
		dbEntry := DBEntry{
			Key:        key,
			Value:      value.(string),
			Expiration: time.Now().Add(24 * time.Hour).Unix(),
		}
		db.Save(&dbEntry)
	}
	return nil
}

// LoadFromDB loads data from the database to the in-memory store.
func LoadFromDB() (map[string]interface{}, error) {
	if db == nil {
		return nil, nil
	}

	var entries []DBEntry
	result := db.Find(&entries)
	if result.Error != nil {
		return nil, result.Error
	}

	data := make(map[string]interface{})
	for _, entry := range entries {
		if entry.Expiration == 0 || entry.Expiration > time.Now().Unix() {
			data[entry.Key] = entry.Value
		}
	}
	return data, nil
}
