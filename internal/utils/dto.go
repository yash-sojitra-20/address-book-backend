package utils

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,strong_password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateContactRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" validate:"omitempty,email"`
	Phone        string `json:"phone" validate:"required"`
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	Country      string `json:"country" validate:"required"`
	Pincode      string `json:"pincode" validate:"required,len=6"`
}


type ContactResponse struct {
	ID			 uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Pincode      string `json:"pincode"`
}