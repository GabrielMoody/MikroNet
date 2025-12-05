package dto

type (
	UserRegistrationsReq struct {
		Email                string `json:"email" validate:"required,email" form:"email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number"`
		Name                 string `json:"name" form:"name" validate:"required"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
	}

	DriverRegistrationsReq struct {
		Name                 string `json:"name" form:"name" validate:"required"`
		Email                string `json:"email" form:"email" validate:"required,email"`
		PhoneNumber          string `json:"phone_number" form:"phone_number"`
		PlateNumber          string `json:"plate_number" form:"plate_number"`
		Password             string `json:"password" form:"password" validate:"required,min=8"`
		PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,eqfield=Password"`
	}

	UserRegistrationsResp struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	UserLoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)
