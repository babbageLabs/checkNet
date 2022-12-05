package utils

import (
	"database/sql"
	"github.com/robfig/cron/v3"
	"log"
	"runtime"
)

var CronJob = cron.New()

func RunMonitor(frequency string) (bool, error) {

	_, err := CronJob.AddFunc(frequency, CollectStatus)
	if err != nil {
		return false, err
	}

	CronJob.Start()

	return true, nil
}

func CollectStatus() {
	db, err := sql.Open("sqlite3", SQLiteDbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	data, err := WifiData(runtime.GOOS)
	if err != nil {
		log.Fatal(err)
	}

	_, err = data.Save(db)
	if err != nil {
		log.Fatal(err)
	}
}

func IsMonitorRunning() bool {
	return len(CronJob.Entries()) > 0
}

func StopMonitor() bool {
	if IsMonitorRunning() {
		CronJob.Stop()
		return true
	}
	return false
}
