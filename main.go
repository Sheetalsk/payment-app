package main

import (
	// Stronger cryptographic hash
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

var db *sql.DB

func main() {
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
	_, err := db.Exec(query) // vulnerable to SQL injection
	if err != nil {
		http.Error(w, "Payment processing error", http.StatusInternalServerError)
		return
	}

	// Weak cryptographic hash (MD5)
	hashedValue := md5.Sum([]byte("sensitiveData"))
	fmt.Fprintf(w, "Processed payment for account %s with amount %s. MD5 hash: %x", accountID, amount, hashedValue)
}
