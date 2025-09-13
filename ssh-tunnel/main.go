package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"ssh-tunnel/sshconfig"
	"ssh-tunnel/tunnel"
)

func main() {
	tun := tunnel.New(
		"ec2-user@jumpbox.us-east-1.mydomain.com",
		sshconfig.PrivateKeyFile("path/to/private/key.pem"),
		"dqrsdfdssdfx.us-east-1.redshift.amazonaws.com:5439",
	)

	tun.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)

	go tun.Start()
	time.Sleep(100 * time.Millisecond)

	connStr := fmt.Sprintf("host=127.0.0.1 port=%d user=foo dbname=mydb sslmode=disable", tun.Local.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open db: %s", err)
	}
	defer db.Close()

	// Тестовий запит
	row := db.QueryRow("SELECT version()")
	var version string
	if err := row.Scan(&version); err != nil {
		log.Fatalf("Query failed: %s", err)
	}
	fmt.Println("Postgres version:", version)
}
