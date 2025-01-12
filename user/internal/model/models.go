package model

type UserDetails struct {
	ID    string `gorm:"column:id;primaryKey;type:varchar(255)" json:"id"`
	Email string `gorm:"column:email;unique;type:varchar(255)" json:"email"`
}

type Review struct {
	ID       int    `gorm:"column:id;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id;type:varchar(255)" json:"user_id"`
	DriverID string `gorm:"column:driver_id;type:varchar(255)" json:"driver_id"`
	Comment  string `gorm:"column:comment;type:varchar(255)" json:"comment"`
	Star     int    `gorm:"column:star" json:"star"`
}
