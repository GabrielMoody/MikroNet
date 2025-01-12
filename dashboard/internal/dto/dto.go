package dto

type (
	GetDriverQuery struct {
		Verified bool `query:"verified"`
	}

	GovRegistrationReq struct {
		FirstName            string `json:"first_name" form:"first_name" validate:"required"`
		LastName             string `json:"last_name" form:"last_name" validate:"required"`
		Email                string `json:"email" form:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number" validate:"required"`
		NIP                  string `json:"nip" form:"nip" validate:"required"`
		Password             string `json:"password" form:"password" validate:"required"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
		ProfilePicture       string `json:"profile_picture" form:"profile_picture" validate:"required"`
	}
)
