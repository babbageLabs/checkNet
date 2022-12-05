package main

import (
	"checkNet/utils"
	"database/sql"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"runtime"
)

func main() {
	app := &cli.App{
		Name:  "checkNet",
		Usage: "Watch and monitor wifi uptime",
		Action: func(*cli.Context) error {
			data, err := utils.WifiData(runtime.GOOS)
			if err != nil {
				return err
			}
			fmt.Printf("your current wifi connection details are \n  %s \n", data.PrettyPrint())
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "monitor",
				Aliases: []string{"m"},
				Usage:   "collect network statistics over time",
				Action: func(cCtx *cli.Context) error {
					db, err := sql.Open("sqlite3", utils.SQLiteDbName)
					if err != nil {
						log.Fatal(err)
					}
					defer func(db *sql.DB) {
						err := db.Close()
						if err != nil {
							log.Fatal(err)
						}
					}(db)

					data, err := utils.WifiData(runtime.GOOS)
					if err != nil {
						return err
					}

					fmt.Printf("Your current wifi connection status \n  %s \n", data.PrettyPrint())

					// save the collected info
					// should it fail for some reason, the user will be aware immediately
					_, err = data.Save(db)
					if err != nil {
						return err
					}

					// now run the scheduled command
					_, err = utils.RunMonitor("0 30 * * * *")
					if err != nil {
						return err
					}

					fmt.Printf("\nNetwork monitoring started. \n" +
						"That is it for now,check back later to see how your network has been performing")
					return nil
				},
			},
			{
				Name:    "stopMonitor",
				Aliases: []string{"sm"},
				Usage:   "stop collection of network statistics over time",
				Action: func(cCtx *cli.Context) error {
					if utils.StopMonitor() {
						fmt.Printf("Your monitor has been stopped successfully, good bye")
					} else {
						fmt.Printf("Your monitor is not running. pleease run 'checkNet m' to start monitoring")
					}
					return nil
				},
			},
			{
				Name:    "inspectMonitor",
				Aliases: []string{"im"},
				Usage:   "check monitoring status",
				Action: func(cCtx *cli.Context) error {
					if utils.IsMonitorRunning() {
						fmt.Printf("Your monitor is running")
					} else {
						fmt.Printf("Your monitor is not running. pleease run 'checkNet m' to start monitoring")
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
