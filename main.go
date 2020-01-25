package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	stdin       *os.File
	connString  string
	queryString string
	warntime    int
	crittime    int
	desiredrows int
)

func EnvOrDefault(env string, defval string) string {
	v, ok := os.LookupEnv(env)
	if ok {
		return v
	}
	return defval
}
func EnvOrDefaultI(env string, defval int) int {
	v, ok := os.LookupEnv(env)
	if ok {
		nv, err := strconv.Atoi(v)
		if err != nil {
			return defval
		}
		return nv
	}
	return defval
}

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-go-mssql-check",
		Short: "The Sensu Go mssql check",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&connString,
		"connstring",
		"C",
		EnvOrDefault("MSSQL_CHECK_CONNSTRING", ""),
		"MSSQL Connection String (MSSQL_CHECK_CONNSTRING)")

	_ = cmd.MarkFlagRequired("connstring")
	cmd.Flags().StringVarP(&queryString,
		"querystring",
		"q",
		EnvOrDefault("MSSQL_CHECK_QUERYSTRING", "select 1"),
		"MSSQL Query String (MSSQL_CHECK_QUERYSTRING)")

	cmd.Flags().IntVarP(&warntime,
		"warntime",
		"w",
		EnvOrDefaultI("MSSQL_CHECK_WARNTIME", -1),
		"MSSQL Query Warning Time (MSSQL_CHECK_WARNTIME)")

	cmd.Flags().IntVarP(&crittime,
		"crittime",
		"c",
		EnvOrDefaultI("MSSQL_CHECK_CRITTIME", -1),
		"MSSQL Query Critical Time (MSSQL_CHECK_CRITTIME)")

	cmd.Flags().IntVarP(&desiredrows,
		"desiredrows",
		"r",
		EnvOrDefaultI("MSSQL_CHECK_DESIREDROWS", -1),
		"MSSQL Query Desired Rows (MSSQL_CHECK_DESIREDROWS)")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		//_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	starttime := time.Now()
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
	start2time := time.Now()
	opentime := start2time.Sub(starttime)
	stmt, err := conn.Prepare(queryString)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	rowcount := 0
	for rows.Next() {
		rowcount++
	}
	// var somenumber int64
	// var somechars string
	// err = rows.Scan(&somenumber, &somechars)
	// if err != nil {
	// 	log.Fatal("Scan failed:", err.Error())
	// }
	endtime := time.Now()
	querytime := endtime.Sub(start2time)
	endtoendtime := endtime.Sub(starttime)
	fmt.Printf("mssql rows=%d,open=%d,query=%d,endtoend=%d %d\n", rowcount, opentime, querytime, endtoendtime, endtime.UnixNano())
	rc := 0
	if desiredrows >= 0 {
		if rowcount != desiredrows {
			fmt.Printf("Row count %d different than desired %d\n", rowcount, desiredrows)
			rc = 2
		}
	}
	if crittime >= 0 {
		if endtoendtime.Nanoseconds() > int64(crittime) {
			fmt.Printf("End to end time %d greater than Critical Threshold %d\n", endtoendtime, crittime)
			rc = 2
		}
	}
	if warntime >= 0 {
		if endtoendtime.Nanoseconds() > int64(warntime) {
			fmt.Printf("End to end time %d greater than Warning Threshold %d\n", endtoendtime, warntime)
			rc = 1
		}
	}
	os.Exit(rc)
	//fmt.Printf("Times: %s %s %s %s %s %d", starttime, start2time, start3time, start4time, endtime, endtime.Sub(starttime))

	// if stdin == nil {
	// 	stdin = os.Stdin
	// }

	// eventJSON, err := ioutil.ReadAll(stdin)
	// if err != nil {
	// 	return fmt.Errorf("failed to read stdin: %s", err)
	// }

	// event := &types.Event{}
	// err = json.Unmarshal(eventJSON, event)
	// if err != nil {
	// 	return fmt.Errorf("failed to unmarshal stdin data: %s", err)
	// }

	// if err = event.Validate(); err != nil {
	// 	return fmt.Errorf("failed to validate event: %s", err)
	// }

	// if !event.HasCheck() {
	// 	return fmt.Errorf("event does not contain check")
	// }

	//return exampleAction(event)
	return nil
}
