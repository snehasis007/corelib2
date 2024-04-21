package base

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Con() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting any
	err = conn.QueryRow(context.Background(), "SELECT first_name FROM employees where employee_id=100").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
