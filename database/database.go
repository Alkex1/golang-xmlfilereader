package database

import (
	"QA-REPORT-EXP/xmlreader"
	"database/sql"
	"fmt"

	// github library
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "alkex"
	dbname   = "QA-automation-reporting"
	password = "Numero123"
	sslmode  = "disable"
)

//DbTest gdfgd
func DbTest(tests []xmlreader.TestCaseDetail) {
	psqlInfo := fmt.Sprintf("host=%s port = %d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlStatement := `
		INSERT INTO testExecutionDetails (
			platformname,
			testcasename, 
			featurename, 
			outcome, 
			startDateTime, 
			endDateTime, 
			duration,
			errorMessage
	)
		Values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)
		RETURNING id`
	for i := 0; i < len(tests); i++ {
		id := 0
		err = db.QueryRow(sqlStatement,
			tests[i].TestSuiteName,
			tests[i].TestCaseName,
			tests[i].FeatureName,
			tests[i].Outcome,
			tests[i].StartDateTime,
			tests[i].EndDateTime,
			tests[i].Duration,
			tests[i].ErroMessage).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success with: ", id)
	}

}
