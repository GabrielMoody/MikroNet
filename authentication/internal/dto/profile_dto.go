package dto

type (
	UserRegistrationsReq struct {
		FirstName            string `json:"first_name" validate:"required"`
		LastName             string `json:"last_name"`
		Email                string `json:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" validate:"required"`
		DateOfBirth          string `json:"date_of_birth" validate:"required"`
		Age                  int    `json:"age" validate:"required"`
		Gender               string `json:"gender" validate:"required"`
		Password             string `json:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}

	UserRegistrationsResp struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
	}

	UserLoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	ForgotPasswordReq struct {
		Email string `json:"email" validate:"required,email"`
	}

	ResetPasswordReq struct {
		Password             string `json:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}
)
