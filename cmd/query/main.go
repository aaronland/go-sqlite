package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-sqlite/v2"
	"github.com/aaronland/go-sqlite/v2/query"
	"log"
	"os"
)

func main() {

	db_uri := flag.String("database-uri", "", "A valid SQLite DSN")
	query_str := flag.String("query", "", "A valid SQL query")

	flag.Parse()

	ctx := context.Background()

	db, err := sqlite.NewDatabase(ctx, *db_uri)

	if err != nil {
		log.Fatalf("Failed to create new database, %v", err)
	}

	conn, err := db.Conn(ctx)

	if err != nil {
		log.Fatalf("Failed to get database connection, %v", err)
	}

	rows, err := conn.Query(*query_str)

	if err != nil {
		log.Fatalf("query err: %s", err)
	}

	wr, err := query.NewCSVQueryWriter(ctx, os.Stdout)

	if err != nil {
		log.Fatalf("Failed to create query writer, %v", err)
	}

	err = query.WriteRows(ctx, wr, rows)

	if err != nil {
		log.Fatalf("Failed to write rows, %v", err)
	}
}
