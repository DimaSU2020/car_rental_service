package users

import (
	"context"
	"errors"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/user/model"
)

var ErrUserNotFound = errors.New("user not found")

type CreateUserInput struct {
	ID        int64     
	Name      string
	Email     string
	Birthday  time.Time
	IsAdmin   bool
	LicenseId *int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateUserInput struct {
	ID        int64     
	Name      string
	Email     string
	Birthday  time.Time
	IsAdmin   bool
	LicenseId *int64
	UpdatedAt time.Time
}

type UserRepository interface {
	List(ctx context.Context, limit, offset int) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

type UserService interface {
	List(ctx context.Context, limit, offset int) ([]*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, input CreateUserInput) (*model.User, error)
	Update(ctx context.Context, input UpdateUserInput) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (u *userService) List(ctx context.Context, limit, offset int) ([]*model.User, error) {
	limitMax := 100
	if limit <= 0 || limit > limitMax { 
		limit = 10
	}
	
	if offset < 0 { 
		offset = 0
	}

	allUsers, err := u.repo.List(ctx, limitMax, offset)
	if err != nil {
		return nil, err
	}

	if offset > len(allUsers) {
		return []*model.User{}, nil
	}

	endPagination := min(limit + offset, len(allUsers))

	if len(allUsers) == 0 {
		return []*model.User{}, nil
	}

	return allUsers[offset:endPagination], nil
}

func (u *userService) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *userService) Create(ctx context.Context, input CreateUserInput) (*model.User, error) {
	user := model.User{
		Name      : input.Name,
		Email     : input.Email,
		Birthday  : input.Birthday,
		IsAdmin   : input.IsAdmin,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return u.repo.Create(ctx, &user)
}

func (u *userService) Update(ctx context.Context, input UpdateUserInput) error {
	user := model.User{
		ID        : input.ID,
		Name      : input.Name,
		Email     : input.Email,
		Birthday  : input.Birthday,
		IsAdmin   : input.IsAdmin,
		LicenseId : input.LicenseId,
		UpdatedAt : time.Now(),
	}

	err := user.Validate()
	if err != nil {
		return err
	}

	return u.repo.Update(ctx, &user)
}

func (u *userService) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}