package dto

type (
	ChangePasswordReq struct {
		OldPassword        string `json:"old_password" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		NewPasswordConfirm string `json:"new_password_confirm" validate:"required,min=8,eqfield=NewPassword"`
	}

	UserRegistrationsReq struct {
		Email                string `json:"email" validate:"required,email" form:"email"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
		Name                 string `json:"name"`
		DateOfBirth          string `json:"date_of_birth"`
		Age                  int    `json:"age"`
	}

	DriverRegistrationsReq struct {
		Name                 string `json:"name" form:"name" validate:"required"`
		Email                string `json:"email" form:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number"`
		SIM                  string `json:"sim" form:"sim"`
		LicenseNumber        string `json:"license_number" form:"license_number"`
		ProfilePicture       string `json:"profile_picture" form:"profile_picture"`
		KTP                  string `json:"ktp" form:"ktp"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
	}

	OwnerRegistrationsReq struct {
		Name                 string `json:"name"`
		Email                string `json:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number"`
		NIK                  string `json:"nik"`
		ProfilePicture       string `json:"profile_picture" form:"profile_picture"`
		Password             string `json:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	}

	GovRegistrationReq struct {
		Name                 string `json:"name" form:"name" validate:"required"`
		Email                string `json:"email" form:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number" validate:"required"`
		NIP                  string `json:"nip" form:"nip" validate:"required"`
		Password             string `json:"password" form:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
		ProfilePicture       string `json:"profile_picture" form:"profile_picture" validate:"required"`
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
