package bussiness

import (
	"SistemVaksinAPI/features/vaksin"
	mocksVaksin "SistemVaksinAPI/features/vaksin/mocks"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	vaksinValue   vaksin.VaksinCore
	vaksinsValue  []vaksin.VaksinCore
	vaksinData    mocksVaksin.Data
	vaksinUseCase vaksin.Bussiness
)

func TestMain(m *testing.M) {
	vaksinUseCase = NewVaksinBussiness(&vaksinData)

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

	os.Exit(m.Run())
}

func TestGetVaksinByID(t *testing.T) {
	t.Run("valid - get vaksin by id", func(t *testing.T) {

		//data := []int{1}
		vaksinData.On("SelectVaksinByID", mock.AnythingOfType("int")).Return(vaksinValue, nil).Once()

		resp, _ := vaksinUseCase.GetVaksinByID(vaksinValue.ID)

		assert.Equal(t, resp.ID, vaksinValue.ID)
	})
}

func TestCreateVaksin(t *testing.T) {
	t.Run("valid - create vaksin", func(t *testing.T) {
		vaksinData.On("SelectVaksin", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, nil).Once()
		vaksinData.On("InsertVaksin", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, nil)

		resp, err := vaksinUseCase.CreateVaksin(vaksinValue)

		assert.Nil(t, err)
		assert.Equal(t, resp.ID, vaksinValue.ID)
	})

	t.Run("invalid - error create donation", func(t *testing.T) {
		vaksinData.On("SelectVaksin", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, nil).Once()
		vaksinData.On("InsertVaksin", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, errors.New("error create vaksin")).Once()

	})
}

func TestUpdateVaksinByID(t *testing.T) {
	t.Run("valid - create vaksin", func(t *testing.T) {
		//vaksinData.On("SelectVaksinByID", mock.AnythingOfType("int")).Return(vaksinValue, nil).Once()
		vaksinData.On("EditVaksinByID", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, nil)

		resp, err := vaksinUseCase.UpdateVaksinByID(vaksin.VaksinCore{
			ID:          1,
			Jenisvaksin: "Astrazenecca",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Jenisvaksin, vaksinValue.Jenisvaksin)
	})

	// t.Run("invalid - error create vaksin", func(t *testing.T) {
	// 	//vaksinData.On("SelectVaksinByID", mock.AnythingOfType("int")).Return(vaksinValue, nil).Once()
	// 	vaksinData.On("EditVaksinByID", mock.AnythingOfType("vaksin.VaksinCore")).Return(vaksinValue, errors.New("error create vaksin")).Once()

	// 	_, err := vaksinUseCase.UpdateVaksinByID(vaksin.VaksinCore{
	// 		Jenisvaksin: "Reboisasi",
	// 	})
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err.Error(), "error edit vaksin")
	// })
}

func TestGetVaksinByFaskesID(t *testing.T) {
	t.Run("valid - get vaksin by faskes id", func(t *testing.T) {
		vaksinData.On("SelectVaksinByFaskesID", mock.AnythingOfType("int")).Return(vaksinData, nil).Once()

		resp, err := vaksinUseCase.GetVaksinByFaskesID(vaksinValue.ID)

		assert.Nil(t, err)
		assert.NotEqual(t, len(resp), 0)
	})
}
