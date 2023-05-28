package controller

import (
	"bitcoin-notify/config"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

type EmailRequest struct {
	Email string `json:"email"`
}

func validateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")

	if !validateEmail(email) {
		http.Error(w, "Invalid email", http.StatusConflict)
		return
	}

	err := saveEmailToFile(email)
	if err != nil {
		log.Println("Error saving email:", err)
		http.Error(w, "Error saving email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Email saved successfully")
}

func saveEmailToFile(email string) error {
	file, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	exists := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == email {
			exists = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("email already exists")
	}

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return err
	}

	return nil
}
