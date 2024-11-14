package database

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"rami/models"

	"github.com/xuri/excelize/v2"
)

type ExcelDB struct {
	filePath     string
	visitorCache []models.Visitor
	logCache     []models.Log
	mutex        sync.RWMutex
}

func NewExcelDB(filePath string) *ExcelDB {
	db := &ExcelDB{filePath: filePath}
	db.reloadData()
	go db.watchFile()
	return db
}

func (db *ExcelDB) GetVisitors() []models.Visitor {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.visitorCache
}

func (db *ExcelDB) GetLogs() []models.Log {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.logCache
}

func (db *ExcelDB) reloadData() {
	db.loadVisitors()
	db.loadLogs()
}

func (db *ExcelDB) loadVisitors() {
	f, err := excelize.OpenFile(db.filePath)
	if err != nil {
		log.Printf("Error opening Excel file: %v", err)
		return
	}
	defer f.Close()

	var visitors []models.Visitor
	rows, err := f.GetRows("Visitors")
	if err != nil {
		log.Printf("Error reading 'Visitors' sheet: %v", err)
		return
	}

	for _, row := range rows[1:] {
		entryApproval, _ := strconv.ParseBool(row[7])
		entryExpiry, _ := strconv.Atoi(row[8])
		clearanceExpiry, _ := strconv.Atoi(row[11])
		securityOfficerApproval, _ := strconv.ParseBool(row[12])
		inside, _ := strconv.ParseBool(row[14])

		visitor := models.Visitor{
			Name:                    row[0],
			CredentialsNumber:       row[1],
			CredentialType:          row[2],
			VehiclePlate:            row[3],
			Association:             row[4],
			Inviter:                 row[5],
			Purpose:                 row[6],
			EntryApproval:           entryApproval,
			EntryExpiry:             entryExpiry,
			SecurityResponse:        row[9],
			ClearanceLevel:          row[10],
			ClearanceExpiry:         clearanceExpiry,
			SecurityOfficerApproval: securityOfficerApproval,
			Notes:                   row[13],
			Inside:                  inside,
		}
		visitors = append(visitors, visitor)
	}

	db.mutex.Lock()
	db.visitorCache = visitors
	db.mutex.Unlock()
	log.Println("Visitor data reloaded.")
}

func (db *ExcelDB) loadLogs() {
	f, err := excelize.OpenFile(db.filePath)
	if err != nil {
		log.Printf("Error opening Excel file: %v", err)
		return
	}
	defer f.Close()

	var logs []models.Log
	rows, err := f.GetRows("Logs")
	if err != nil {
		log.Printf("Error reading 'Logs' sheet: %v", err)
		return
	}

	for _, row := range rows[1:] {
		logEntry := models.Log{
			Event:     row[0],
			Serial:    row[1],
			Timestamp: row[2],
		}
		logs = append(logs, logEntry)
	}

	db.mutex.Lock()
	db.logCache = logs
	db.mutex.Unlock()
	log.Println("Log data reloaded.")
}

func (db *ExcelDB) watchFile() {
	lastModified := time.Now()
	for {
		time.Sleep(2 * time.Second)
		info, err := os.Stat(db.filePath)
		if err != nil {
			log.Printf("Error watching file: %v", err)
			continue
		}
		if info.ModTime().After(lastModified) {
			lastModified = info.ModTime()
			db.reloadData()
		}
	}
}
