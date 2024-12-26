package dto

type (
	OwnerRegistrationReq struct {
		FirstName       string `json:"first_name" validate:"required"`
		LastName        string `json:"last_name" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		PhoneNumber     string `json:"phone_number" validate:"required"`
		NIK             string `json:"nik" validate:"required"`
		Mikrolet        string `json:"mikrolet" validate:"required"`
		Password        string `json:"password" validate:"required,min=8"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}

	DriverRegistrationReq struct {
		FirstName          string `json:"first_name" validate:"required"`
		LastName           string `json:"last_name" validate:"required"`
		PhoneNumber        string `json:"phone_number" validate:"required"`
		DateOfBirth        string `json:"date_of_birth" validate:"required"`
		Age                int    `json:"age" validate:"required,number"`
		RegistrationNumber string `json:"registration_number" validate:"required,eqfield=RegistrationNumber"`
		Password           string `json:"password" validate:"required,min=8"`
		PasswordConfirm    string `json:"password_confirm" validate:"required,eqfield=Password"`
	}
)
