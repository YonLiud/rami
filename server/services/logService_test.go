package services

import (
	"rami/models"
	"testing"
)

func GenerateRandomLog() models.Log {
	return models.Log{
		Serial:    generateRandomString(10),                        // Random Serial (e.g., "mG522mmhPb")
		Event:     generateRandomChoice([]string{"Entry", "Exit"}), // Random Event (either "Entry" or "Exit")
		Timestamp: generateRandomTimestamp(),                       // Random Timestamp
	}
}

func GenerateRandomLogForSerial(serial string) models.Log {
	return models.Log{
		Serial:    serial,                                          // Fixed Serial
		Event:     generateRandomChoice([]string{"Entry", "Exit"}), // Random Event (either "Entry" or "Exit")
		Timestamp: generateRandomTimestamp(),                       // Random Timestamp
	}
}

func TestCreateLog(t *testing.T) {
	InitiateTestDB()
	demoLog := GenerateRandomLog()

	err := CreateLog(&demoLog)
	if err != nil {
		t.Errorf("Error creating log: %v", err)
	}

	logs, err := GetLogsBySerial(demoLog.Serial)
	if err != nil {
		t.Errorf("Error retrieving log: %v", err)
	}

	log := logs[0]

	if log.Serial != demoLog.Serial {
		t.Errorf("Expected serial %s, got %s", demoLog.Serial, log.Serial)
	}
	if log.Event != demoLog.Event {
		t.Errorf("Expected event %s, got %s", demoLog.Event, log.Event)
	}
	if log.Timestamp != demoLog.Timestamp {
		t.Errorf("Expected timestamp %d, got %d", demoLog.Timestamp, log.Timestamp)
	}
}

func TestGetAllLogs(t *testing.T) {
	InitiateTestDB()

	// array of 5 random logs
	var demoLogs []models.Log
	for i := 0; i < 5; i++ {
		demoLogs = append(demoLogs, GenerateRandomLog())
	}

	// create each log
	for i := 0; i < len(demoLogs); i++ {
		err := CreateLog(&demoLogs[i])
		if err != nil {
			t.Errorf("Error creating log: %v", err)
		}
	}

	logs, err := GetAllLogs()
	if err != nil {
		t.Errorf("Error retrieving logs: %v", err)
	}

	if len(logs) != len(demoLogs) {
		t.Errorf("Expected %d logs, got %d", len(demoLogs), len(logs))
	}

	for i := 0; i < len(demoLogs); i++ {
		if logs[i].Serial != demoLogs[i].Serial {
			t.Errorf("Expected serial %s, got %s", demoLogs[i].Serial, logs[i].Serial)
		}
		if logs[i].Event != demoLogs[i].Event {
			t.Errorf("Expected event %s, got %s", demoLogs[i].Event, logs[i].Event)
		}
		if logs[i].Timestamp != demoLogs[i].Timestamp {
			t.Errorf("Expected timestamp %d, got %d", demoLogs[i].Timestamp, logs[i].Timestamp)
		}
	}
}

func TestGetLogsBySerial(t *testing.T) {
	InitiateTestDB()

	randomSerial := generateRandomString(10)

	// array of 10 random logs
	var demoLogs []models.Log
	for i := 0; i < 10; i++ {
		demoLogs = append(demoLogs, GenerateRandomLog())
	}

	// pick 3 random logs and set their serial to randomSerial
	serialSet := make(map[int]struct{})
	for len(serialSet) < 3 { // Ensure we select 3 unique logs
		randomInt := generateRandomInt(0, len(demoLogs)-1) // Ensure within bounds
		serialSet[randomInt] = struct{}{}                  // Add index to the set (unique)
	}

	// Assign the serial to the selected logs
	for index := range serialSet {
		demoLogs[index].Serial = randomSerial
	}

	for i := 0; i < len(demoLogs); i++ {
		err := CreateLog(&demoLogs[i])
		if err != nil {
			t.Errorf("Error creating log %d: %v", i, err)
		}
	}

	logs, err := GetLogsBySerial(randomSerial)
	if err != nil {
		t.Errorf("Error retrieving logs: %v", err)
	}

	if len(logs) != 3 {
		t.Errorf("Expected 3 logs, got %d", len(logs))
		t.FailNow()
	}

	for i := 0; i < 3; i++ {
		if logs[i].Serial != randomSerial {
			t.Errorf("Expected serial %s, got %s", randomSerial, logs[i].Serial)
		}
	}
}
