package db

import (
	"context"
	"fmt"
	"log"

	
	"github.com/jackc/pgx/v5"
)

func ConnectToSupabase() *pgx.Conn {
	url := "postgresql://postgres:CloudDB_2025!@db.jkvisgpgreahkkhpoyal.supabase.co:5432/postgres"

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to Supabase: %v\n", err)
	}

	fmt.Println("✅ Connected to Supabase")
	return conn
}

