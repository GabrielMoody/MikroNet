package dto

type (
	UserRegistrationsReq struct {
		NamaLengkap         string `json:"nama_lengkap" form:"nama_lengkap" validate:"required"`
		Email               string `json:"email" form:"email" validate:"required,email"`
		NomorTelepon        string `json:"nomor_telepon" form:"nomor_telepon" validate:"required"`
		KataSandi           string `json:"kata_sandi" form:"kata_sandi" validate:"required,min=8"`
		KonfirmasiKataSandi string `json:"konfirmasi_kata_sandi" form:"konfirmasi_kata_sandi" validate:"required,eqfield=KataSandi"`
	}

	UserRegistrationsResp struct {
		ID           string `json:"id"`
		NamaLengkap  string `json:"nama_lengkap"`
		Email        string `json:"email"`
		NomorTelepon string `json:"nomor_telepon"`
	}

	UserLoginReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserChangeProfileReq struct {
		NamaLengkap  string `json:"nama_lengkap,omitempty" form:"nama_lengkap"`
		Email        string `json:"email,omitempty" form:"email" validate:"email"`
		NomorTelepon string `json:"nomor_telepon,omitempty" form:"nomor_telepon"`
		JenisKelamin string `json:"jenis_kelamin,omitempty" form:"jenis_kelamin"`
	}

	ChangePasswordReq struct {
		OldPassword        string `json:"old_password" form:"old_password"`
		NewPassword        string `json:"new_password" form:"new_password" validate:"required,min=8"`
		NewPasswordConfirm string `json:"new_password_confirm" form:"new_password_confirm" validate:"required,min=8,eqfield=NewPassword"`
	}
)
