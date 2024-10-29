package dto

type (
	UserRegistrationsResp struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Role        string `json:"role"`
		ImageUrl    string `json:"image_url"`
	}

	UserChangeProfileReq struct {
		FirstName   string `form:"first_name" json:"first_name,omitempty"`
		LastName    string `form:"last_name" json:"last_name,omitempty"`
		Email       string `form:"email,omitempty" json:"email,omitempty" validate:"omitempty,email"`
		Gender      string `form:"gender" json:"gender,omitempty"`
		DateOfBirth string `form:"date_of_birth" json:"date_of_birth,omitempty"`
		Age         int    `form:"age" json:"age,omitempty"`
		Image       string `form:"image,omitempty"`
	}

	ChangePasswordReq struct {
		OldPassword        string `json:"old_password"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		NewPasswordConfirm string `json:"new_password_confirm" validate:"required,min=8,eqfield=NewPassword"`
	}
)
