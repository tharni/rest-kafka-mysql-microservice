package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func saveJobsToDb(job Job) {
	db, err := sql.Open("mysql", "root:admin@(127.0.0.1:3306)/edb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Insert a new job

		result, err := db.Exec(`INSERT INTO jobs (title, description, company, salary) VALUES (?, ?, ?, ?)`, job.Title, job.Description, job.Company, job.Salary)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{ // Query all users

		rows, err := db.Query(`SELECT id, title, description, company, salary FROM jobs`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var jobs []Job
		for rows.Next() {
			var job Job

			err := rows.Scan(&job.id, &job.Title, &job.Description, &job.Company, &job.Salary)
			if err != nil {
				log.Fatal(err)
			}
			jobs = append(jobs, job)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", jobs)
	}

}
