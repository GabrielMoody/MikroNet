package models

type BlockedAccount struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	AccountID string `gorm:"column:account_id;unique;type:varchar(255)" json:"account_id"`
	Role      string `gorm:"column:role;type:varchar(255)" json:"role"`
}

type GovDetails struct {
	ID             string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	Name           string `gorm:"column:name;type:varchar(255)" json:"name"`
	Email          string `gorm:"column:email;type:varchar(255)" json:"email"`
	NIP            string `gorm:"column:nip;type:varchar(255)" json:"nip"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture"`
}

type OwnerDetails struct {
	ID             string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	Name           string `gorm:"column:name;type:varchar(255)" json:"name"`
	Email          string `gorm:"column:email;type:varchar(255)" json:"email"`
	PhoneNumber    string `gorm:"column:phone_number;type:varchar(255)" json:"phone_number"`
	NIK            string `gorm:"column:nik;type:varchar(255)" json:"nik"`
	Verified       bool   `gorm:"column:verified;default:false" json:"verified"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture"`
}

type Admin struct {
	ID   string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}
