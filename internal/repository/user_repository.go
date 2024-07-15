package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return db.Where("token = ?", token).First(user).Error
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) CreateVerificationToken(db *gorm.DB, user *entity.User, token string) error {
	verificationToken := entity.VerificationToken{
		UserId: user.ID,
		Token:  token,
	}
	return db.Table(verificationToken.TableName()).Create(&verificationToken).Error
}

func (r *UserRepository) UpdateVerificationToken(db *gorm.DB, user *entity.User, token string) error {
	verificationToken := entity.VerificationToken{
		UserId: user.ID,
		Token:  token,
	}
	return db.Table(verificationToken.TableName()).Save(&verificationToken).Error
}

func (r *UserRepository) VerifyEmailByToken(db *gorm.DB, user *entity.User, token string) error {
	verificationToken := new(entity.VerificationToken)
	return db.Table(verificationToken.TableName()).Where("token = ?", token).First(user).Error
}
