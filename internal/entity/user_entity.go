package entity

// User is a struct that represents a user entity
type User struct {
	ID         string    `gorm:"column:id;primaryKey"`
	Password   string    `gorm:"column:password"`
	Name       string    `gorm:"column:name"`
	Email      string    `gorm:"column:email"`
	Token      string    `gorm:"column:token"`
	CreatedAt  int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	VerifiedAt int64     `gorm:"column:verified_at"`
	Contacts   []Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
