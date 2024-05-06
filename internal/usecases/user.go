package usecases

import (
	"github.com/cleyton1986/client-cleancode-solid/internal/entities"
	"github.com/cleyton1986/client-cleancode-solid/internal/interfaces/repositories"
)

type UserRepository = *repositories.UserRepository  // Mudança aqui para aceitar ponteiro

type UserUseCase struct {
    repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
    return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(user *entities.User) error {
    return uc.repo.Create(user)
}

func (uc *UserUseCase) ListUsers() ([]entities.User, error) {
    return uc.repo.FindAll()
}

func (uc *UserUseCase) UpdateUser(user *entities.User) error {
    return uc.repo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id uint) error {
    return uc.repo.Delete(id)
}
