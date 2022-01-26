package bussiness

import (
	"SistemVaksinAPI/features/user"
	mocksUser "SistemVaksinAPI/features/user/mocks"
	"os"
	"testing"
)

var (
	userValue   user.UserCore
	usersValue  []user.UserCore
	userUseCase user.Bussiness
	userData    mocksUser.Data
)

func TestMain(m *testing.M) {
	userUseCase = NewUserBussiness(&userData, &fas)

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
