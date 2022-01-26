package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	mocksFaskes "SistemVaksinAPI/features/faskes/mocks"
	"SistemVaksinAPI/features/requestvaksin"
	mocksRequestvaksin "SistemVaksinAPI/features/requestvaksin/mocks"
	"SistemVaksinAPI/features/user"
	mocksUser "SistemVaksinAPI/features/user/mocks"
	"SistemVaksinAPI/features/vaksin"
	mocksVaksin "SistemVaksinAPI/features/vaksin/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userValue           user.UserCore
	usersValue          []user.UserCore
	userUseCase         user.Bussiness
	userData            mocksUser.Data
	vaksinValue         vaksin.VaksinCore
	vaksinsValue        []vaksin.VaksinCore
	vaksinData          mocksVaksin.Bussiness
	requestvaksinData   mocksRequestvaksin.Bussiness
	faskesData          mocksFaskes.Bussiness
	faskesValue         faskes.FaskesCore
	faskessValue        []faskes.FaskesCore
	requestvaksinValue  requestvaksin.RequestvaksinCore
	requestvaksinsValue []requestvaksin.RequestvaksinCore
)

func TestMain(m *testing.M) {
	userUseCase = NewUserBussiness(&userData, &vaksinData, &faskesData, &requestvaksinData)

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

	vaksinValue = vaksin.VaksinCore{
		ID:          1,
		Jenisvaksin: "Sinovac",
		Jadwal:      "17 April 2021",
		Waktu:       "08.00-09.00",
		FaskesID:    1,
	}

	vaksinsValue = []vaksin.VaksinCore{
		{
			ID:          1,
			Jenisvaksin: "Sinovac",
			Jadwal:      "17 April 2021",
			Waktu:       "08.00-09.00",
			FaskesID:    1,
		},
	}

	faskesValue = faskes.FaskesCore{
		ID:   1,
		Nama: "Rumah Sakit Sumber Waraas",
	}

	faskessValue = []faskes.FaskesCore{
		{
			ID:   1,
			Nama: "Rumah Sakit Sumber Waraas",
		},
	}

	requestvaksinValue = requestvaksin.RequestvaksinCore{
		ID:            1,
		Nama:          "Djoko Anwar",
		NIK:           214145,
		JenisKelamin:  "Laki Laki",
		TanggalLahir:  "17 Januari 2021",
		Nomor:         "014",
		UserID:        1,
		VaksinID_satu: 2,
		Status_satu:   "Sudah",
		VaksinID_dua:  2,
		Status_dua:    "Sudah",
	}

	requestvaksinsValue = []requestvaksin.RequestvaksinCore{
		{
			ID:            1,
			Nama:          "Djoko Anwar",
			NIK:           214145,
			JenisKelamin:  "Laki Laki",
			TanggalLahir:  "17 Januari 2021",
			Nomor:         "014",
			UserID:        1,
			VaksinID_satu: 2,
			Status_satu:   "Sudah",
			VaksinID_dua:  2,
			Status_dua:    "Sudah",
		},
	}

	os.Exit(m.Run())
}

func TestGetUserByID(t *testing.T) {
	t.Run("valid - get user by id", func(t *testing.T) {
		userData.On("SelectUserByID", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		vaksinData.On("GetVaksinByID", mock.AnythingOfType("int")).Return(vaksinsValue[0], nil).Twice()
		faskesData.On("GetFaskesByID", mock.AnythingOfType("int")).Return(faskessValue[0], nil).Twice()
		requestvaksinData.On("GetRequestvaksinByUserID", mock.AnythingOfType("int")).Return(requestvaksinsValue, nil).Twice()

		resp, _ := userUseCase.GetUserByID(usersValue[0].UserID)

		assert.Equal(t, resp.UserID, usersValue[0].UserID)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("valid - create user", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		userData.On("InsertUser", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		_, err := userUseCase.CreateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.Nil(t, err)
	})

	t.Run("invalid - error create user", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		userData.On("InsertUser", mock.AnythingOfType("user.UserCore")).Return(userValue, errors.New("error insert user")).Once()
		_, err := userUseCase.CreateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error insert user")
	})

	t.Run("invalid email - data is available", func(t *testing.T) {
		userData.On("SelectUserEmail", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		_, err := userUseCase.CreateUser(userValue)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "data is available")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("valid - update user", func(t *testing.T) {
		userData.On("SelectUserByID", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		userData.On("EditUser", mock.AnythingOfType("user.UserCore")).Return(userValue, nil).Once()
		resp, err := userUseCase.UpdateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Email, userValue.Email)
	})

	t.Run("invalid - error update user", func(t *testing.T) {
		userData.On("SelectUserByID", mock.AnythingOfType("int")).Return(usersValue[0], nil).Once()
		userData.On("EditUser", mock.AnythingOfType("user.UserCore")).Return(userValue, errors.New("error edit user")).Once()
		_, err := userUseCase.UpdateUser(user.UserCore{
			Email: "barrykeoghan2@gmail.com",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error edit user")
	})
}
