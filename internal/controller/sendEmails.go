package controller

import (
	"bitcoin-notify/config"
	"bufio"
	"log"
	"net/http"
	"os"
)

func SendEmails(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(config.FilePath)
	if err != nil {
		log.Println("Error reading email:", err)
		http.Error(w, "Error reading email", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var emails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading email:", err)
		http.Error(w, "Error reading email", http.StatusInternalServerError)
		return
	}

}
