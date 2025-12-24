package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
)

func GenerateContactsCSV(userID uint, contacts []models.Contact) (string, error) {

	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf(
		"contacts_user_%d_%s.csv",
		userID,
		timestamp,
	)

	file, err := os.Create(fileName)
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

	return fileName, nil
}
