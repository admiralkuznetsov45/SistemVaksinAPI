package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	"SistemVaksinAPI/features/requestvaksin"
	"SistemVaksinAPI/features/user"
	mocks "SistemVaksinAPI/features/user/mocks"
	"SistemVaksinAPI/features/vaksin"
	"os"
	"testing"
)

var (
	userValue         user.UserCore
	usersValue        []user.UserCore
	userUseCase       user.Bussiness
	userData          mocks.Data
	vaksinUsecase     vaksin.Bussiness
	faskesData        faskes.Bussiness
	requestvaksinData requestvaksin.Bussiness
)

func TestMain(m *testing.M) {
	userUseCase = NewUserBussiness(&userData)

	userValue = user.UserCore{
		UserID:      1,
		Namalengkap: "Barry Keoghan",
		Email:       "barrykeoghan1@gmail.com",
		Password:    "barrykeoghan1",
	}

	usersValue = []user.UserCore{
		{
			UserID:      1,
			Namalengkap: "Barry Keoghan",
			Email:       "barrykeoghan1@gmail.com",
			Password:    "barrykeoghan1",
		},
	}

	os.Exit(m.Run())
}
