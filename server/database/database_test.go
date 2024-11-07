package database

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	db := InitDB("test.db")

	if db == nil {
		t.Fatalf("Expected database connection to be established, got nil")
	}

	if err := db.Exec("SELECT 1").Error; err != nil {
		t.Fatalf("Failed to execute test query: %v", err)
	}

	if !db.Migrator().HasTable("visitors") {
		t.Errorf("Expected 'visitors' table to be created, but it wasn't")
	}
	if !db.Migrator().HasTable("logs") {
		t.Errorf("Expected 'logs' table to be created, but it wasn't")
	}
	if !db.Migrator().HasTable("csos") {
		t.Errorf("Expected 'csos' table to be created, but it wasn't")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get underlying database connection: %v", err)
	}
	sqlDB.Close()

	os.Remove("test.db")
	if _, err := os.Stat("test.db"); err == nil {
		t.Errorf("Expected database file to be removed, but it wasn't")
	}

}

func TestGetDB(t *testing.T) {
	db := InitDB("test.db")

	if db != GetDB() {
		t.Errorf("Expected GetDB() to return the same database connection, but it didn't")
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get underlying database connection: %v", err)
	}
	sqlDB.Close()

	os.Remove("test.db")
	if _, err := os.Stat("test.db"); err == nil {
		t.Errorf("Expected database file to be removed, but it wasn't")
	}

}
