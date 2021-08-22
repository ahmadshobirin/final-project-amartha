package user_test

import (
	"context"
	"main-backend/bussiness/role"
	"main-backend/bussiness/user"
	userMock "main-backend/bussiness/user/mocks"
	"main-backend/helper/messages"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRepo    userMock.Repository
	userUsecase user.Usecase
)

func setup() {
	userUsecase = user.NewUserUsecase(2, &userRepo)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := []user.Domain{
			{
				ID:     1,
				RoleID: 1,
				Role: &role.Domain{
					ID:   1,
					Code: "AM",
					Name: "Admin",
				},
				Name:  "Ahmad Shobirin",
				Email: "ahmadshobirin@gmail.com",
			},
			{
				ID:     1,
				RoleID: 1,
				Role: &role.Domain{
					ID:   1,
					Code: "AM",
					Name: "Admin",
				},
				Name:  "Ahmad Shobirin",
				Email: "ahmadshobirin@gmail.com",
			},
		}

		userRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, 0, nil).Once()

		result, total, err := userUsecase.Fetch(context.Background(), "AM", 1, 10)

		assert.Nil(t, err)
		assert.GreaterOrEqual(t, 0, total)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 2, invalid param ", func(t *testing.T) {
		domain := []user.Domain{
			{
				ID:     1,
				RoleID: 1,
				Role: &role.Domain{
					ID:   1,
					Code: "AM",
					Name: "Admin",
				},
				Name:  "Ahmad Shobirin",
				Email: "ahmadshobirin@gmail.com",
			},
			{
				ID:     1,
				RoleID: 1,
				Role: &role.Domain{
					ID:   1,
					Code: "AM",
					Name: "Admin",
				},
				Name:  "Ahmad Shobirin",
				Email: "ahmadshobirin@gmail.com",
			},
		}

		userRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, 0, nil).Once()

		result, total, err := userUsecase.Fetch(context.Background(), "AM", -1, -10)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, 0, total)
		assert.Equal(t, 2, len(result))
	})

	t.Run("test case 3, repo err", func(t *testing.T) {

		userRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]user.Domain{}, 0, gorm.ErrRecordNotFound).Once()

		result, count, err := userUsecase.Fetch(context.Background(), "AM", -1, -10)

		assert.Equal(t, result, []user.Domain{})
		assert.Equal(t, count, 0)
		assert.Equal(t, err, gorm.ErrRecordNotFound)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := user.Domain{
			ID:     1,
			RoleID: 1,
			Name:   "Ahmad Shobirin",
			Email:  "ahmadshobirin@gmail.com",
		}

		userRepo.On("FindByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := userUsecase.FindByID(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := userUsecase.FindByID(context.Background(), -1)

		assert.Equal(t, result, user.Domain{})
		assert.Equal(t, err, messages.ErrIDNotFound)
	})

	t.Run("test case 3, repo err", func(t *testing.T) {
		userRepo.On("FindByID", mock.Anything, mock.AnythingOfType("string")).Return(user.Domain{}, messages.ErrInvalidParam).Once()
		result, err := userUsecase.FindByEmail(context.Background(), "")

		assert.Equal(t, result, user.Domain{})
		assert.Equal(t, err, messages.ErrInvalidParam)
	})
}

func TestFindByEmail(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := user.Domain{
			ID:     1,
			RoleID: 1,
			Name:   "Ahmad Shobirin",
			Email:  "ahmadshobirin@gmail.com",
		}

		userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := userUsecase.FindByEmail(context.Background(), "ahmadshobirin@gmail.com")
		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, repo err", func(t *testing.T) {
		userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(user.Domain{}, messages.ErrInvalidParam).Once()
		result, err := userUsecase.FindByEmail(context.Background(), "")

		assert.Equal(t, result, user.Domain{})
		assert.Equal(t, err, messages.ErrInvalidParam)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := user.Domain{
			ID:     1,
			RoleID: 1,
			Name:   "Ahmad Shobirin",
			Email:  "ahmadshobirin@gmail.com",
		}

		userRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(user.Domain{}, nil).Once()
		userRepo.On("Store", mock.Anything, mock.AnythingOfType("*user.Domain")).Return(user.Domain{}, nil).Once()

		_, err := userUsecase.Store(context.Background(), &domain, 1)

		assert.NoError(t, err)
	})
}
