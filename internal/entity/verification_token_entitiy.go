package entity

type VerificationToken struct {
	ID        string `gorm:"column:id;primaryKey"`
	Token     string `gorm:"column:token"`
	UserId    string `gorm:"column:user_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User      User   `gorm:"foreignKey:user_id;references:id"`
}

func (c *VerificationToken) TableName() string {
	return "user_verification_tokens"
}
