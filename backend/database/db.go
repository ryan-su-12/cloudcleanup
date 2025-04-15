package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectToSupabase() *pgx.Conn {
	url := "postgres://your_user:your_password@your_host:5432/your_database"

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to Supabase: %v\n", err)
	}

	fmt.Println("✅ Connected to Supabase")
	return conn
}

