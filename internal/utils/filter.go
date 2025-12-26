package utils

import "github.com/yash-sojitra-20/address-book-backend/internal/models"

func FilterAddressFields(addresses []models.Address, fields []string) [][]string {
	var rows [][]string

	// Add header row (the selected fields)
	rows = append(rows, fields)

	for _, a := range addresses {
		var row []string

		for _, f := range fields {
			switch f {
			case "first_name":
				row = append(row, a.FirstName)
			case "last_name":
				row = append(row, a.LastName)
			case "email":
				row = append(row, a.Email)
			case "phone":
				row = append(row, a.Phone)
			case "address_line1":
				row = append(row, a.AddressLine1)
			case "address_line2":
				row = append(row, a.AddressLine2)
			case "city":
				row = append(row, a.City)
			case "state":
				row = append(row, a.State)
			case "country":
				row = append(row, a.Country)
			case "pincode":
				row = append(row, a.Pincode)
			}
		}

		rows = append(rows, row)
	}

	return rows
}
