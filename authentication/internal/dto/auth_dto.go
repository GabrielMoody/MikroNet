package dto

type (
	ChangePasswordReq struct {
		OldPassword        string `json:"old_password" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		NewPasswordConfirm string `json:"new_password_confirm" validate:"required,min=8,eqfield=NewPassword"`
	}

	UserRegistrationsReq struct {
		FirstName            string `json:"first_name" form:"first_name"`
		LastName             string `json:"last_name" form:"last_name"`
		Email                string `json:"email" validate:"required,email" form:"email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number"`
		DateOfBirth          string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
		Age                  int    `json:"age" form:"age"`
		ProfilePicture       string `json:"profile_picture" form:"profile_picture" form:"profile_picture"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
	}

	DriverRegistrationsReq struct {
		FirstName            string `json:"first_name" form:"first_name" validate:"required"`
		LastName             string `json:"last_name" form:"last_name" `
		Email                string `json:"email" form:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number"`
		DateOfBirth          string `json:"date_of_birth" form:"date_of_birth"`
		Age                  int    `json:"age" form:"age"`
		LicenseNumber        string `json:"license_number" form:"license_number" `
		ProfilePicture       string `json:"profile_picture" form:"profile_picture" validate:"required"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
	}

	OwnerRegistrationsReq struct {
		FirstName            string `json:"first_name"`
		LastName             string `json:"last_name"`
		Email                string `json:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number"`
		NIK                  string `json:"nik"`
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
