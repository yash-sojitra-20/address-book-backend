package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
)

func GenerateContactsCSV(userID uint, contacts []models.Contact) (string, error) {

	timestamp := time.Now().Format("20060102_150405")

	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(baseDir, "exports")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf(
		"%s/contacts_user_%d_%s.csv",
		dir,
		userID,
		timestamp,
	)

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"First Name",
		"Last Name",
		"Email",
		"Phone",
		"City",
		"State",
		"Country",
		"Pincode",
	})

	// Data
	for _, c := range contacts {
		writer.Write([]string{
			c.FirstName,
			c.LastName,
			c.Email,
			c.Phone,
			c.City,
			c.State,
			c.Country,
			c.Pincode,
		})
	}

	return filePath, nil
}
