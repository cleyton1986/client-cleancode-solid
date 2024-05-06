package repositories

import (
	"github.com/cleyton1986/client-cleancode-solid/internal/entities"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (repo *UserRepository) Create(user *entities.User) error {
    return repo.db.Create(user).Error
}

func (repo *UserRepository) FindAll() ([]entities.User, error) {
    var users []entities.User
    err := repo.db.Find(&users).Error
    return users, err
}

func (repo *UserRepository) Update(user *entities.User) error {
    return repo.db.Save(user).Error
}

func (repo *UserRepository) Delete(id uint) error {
    return repo.db.Where("id = ?", id).Delete(&entities.User{}).Error
}
