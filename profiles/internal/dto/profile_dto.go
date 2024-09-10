package dto

type (
	UserRegistrationsResp struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
	}

	UserChangeProfileReq struct {
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty" validate:"email"`
		Gender      string `json:"gender,omitempty"`
		DateOfBirth string `json:"date_of_birth,omitempty"`
		Age         int    `json:"age,omitempty"`
		ImageURL    string `json:"image_url,omitempty"`
	}

	ChangePasswordReq struct {
		OldPassword        string `json:"old_password"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		NewPasswordConfirm string `json:"new_password_confirm" validate:"required,min=8,eqfield=NewPassword"`
	}
)
