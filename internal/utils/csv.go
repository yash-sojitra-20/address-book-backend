package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yash-sojitra-20/address-book-backend/internal/models"
)

func GenerateAddressesCSV(userID uint, addresses []models.Address) (string, error) {

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
		"%s/addresses_user_%d_%s.csv",
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
	for _, c := range addresses {
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

func GenerateCustomAddressesCSV(userID uint, rows [][]string) (string, string, error) {

	timestamp := time.Now().Format("20060102_150405")

	baseDir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}

	dir := filepath.Join(baseDir, "exports")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", "", err
	}

	// File name for custom export
	// filePath := fmt.Sprintf(
	// 	"%s/address_custom_%d_%s.csv",
	// 	dir,
	// 	userID,
	// 	timestamp,
	// )

	fileName := fmt.Sprintf(
		"address_custom_%d_%s.csv",
		userID,
		timestamp,
	)

	filePath := filepath.Join(dir, fileName)

	// Create file
	file, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write all rows (header + data)
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return "", "", err
		}
	}

	return filePath, fileName, nil
}
