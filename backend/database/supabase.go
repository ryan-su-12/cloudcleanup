package database

import (
	"context"
	"fmt"
	"log"
	"os" // 👈 Add this!

	"github.com/jackc/pgx/v5"
)

var SupabaseConn *pgx.Conn

func ConnectToSupabase() {
	// Use your full connection string here
	url := os.Getenv("SUPABASE_URL")

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("❌ Unable to connect to Supabase: %v", err)
	}

	fmt.Println("✅ Connected to Supabase")
	SupabaseConn = conn
}
