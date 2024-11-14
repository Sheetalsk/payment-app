package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"crypto/md5"       // Weak cryptographic hash
	"crypto/sha1"      // Weak cryptographic hash
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Hardcoded database credentials
	dbUser := "admin"
	dbPass := "password123"
	dbName := "paymentDB"
	
	// Initialize database connection
	connStr := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Start the HTTP server
	http.HandleFunc("/pay", paymentHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Vulnerable payment handler
func paymentHandler(w http.ResponseWriter, r *http.Request) {
	// Get user input without validation
	accountID := r.URL.Query().Get("accountID")
	amount := r.URL.Query().Get("amount")

	// SQL Injection vulnerability
	query := fmt.Sprintf("UPDATE accounts SET balance = balance - %s WHERE id = '%s'", amount, accountID)
	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, "Payment processing error", http.StatusInternalServerError)
		return
	}

	// Weak cryptographic hash
	hashedPassword := md5.Sum([]byte("sensitivePassword"))
	fmt.Fprintf(w, "Payment processed for account %s with amount %s. Hashed password (MD5): %x", accountID, amount, hashedPassword)
}

func insecureHashingExample() {
	// Another insecure hash example
	data := "sensitiveInformation"
	hashedData := sha1.Sum([]byte(data))
	fmt.Printf("SHA1 hash of sensitive data: %x\n", hashedData)
}

func logSensitiveInfo() {
	// Exposing sensitive information in logs
	userToken := "user-secret-token"
	fmt.Println("User token:", userToken)
}
