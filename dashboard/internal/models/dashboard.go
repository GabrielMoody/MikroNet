package models

type BlockedAccount struct {
	ID        int    `gorm:"column:id;primaryKey" json:"id"`
	AccountID string `gorm:"column:account_id;unique" json:"account_id"`
	Role      string `gorm:"column:role" json:"role"`
}

type GovDetails struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	NIP       string `gorm:"column:nip" json:"nip"`
}

type OwnerDetails struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	NIK       string `gorm:"column:nik" json:"nik"`
	Verified  bool   `gorm:"column:verified" json:"verified"`
}

type Admin struct {
	ID        string `gorm:"column:id;primaryKey" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
}
