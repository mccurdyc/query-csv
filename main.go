package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"

	// DB drivers
	_ "github.com/lib/pq"           // postgres
	_ "github.com/mattn/go-sqlite3" // sqlite
)

// Run an arbitrary query where the parameter markers in the query should line up
// in order with the values in each line of the CSV file.
//
// Also, don't include a header line in the CSV file.
func main() {
	var (
		file    = flag.String("file", "", "CSV file with data")
		connStr = flag.String("conn_str", "", "database connection string")
		driver  = flag.String("db_driver", "postgres", "the database driver")
		query   = flag.String("query", "", "the query to execute")
	)

	flag.Parse()

	db, err := sql.Open(*driver, *connStr)
	if err != nil {
		fmt.Println(errors.Wrap(err, "couldn't connect to database"))
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println(errors.Wrap(err, "couldn't open file"))
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(f))

	i := 1
	for {
		fmt.Println(i)
		line, err := reader.Read()

		// we have to convert []string -> []interface for db.Exec, below
		vals := make([]interface{}, len(line))
		for i, v := range line {
			vals[i] = v
		}

		switch err {
		case nil:
			_, err := db.Exec(*query, vals...)
			if err != nil {
				fmt.Println(errors.Wrap(err, "query failed"))
			}
		case io.EOF:
			fmt.Println("finished")
			os.Exit(0)
		default:
			fmt.Println(errors.Wrap(err, "error reading line from CSV file"))
			continue
		}
		i++
	}
}
