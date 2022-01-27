package bussiness

import (
	"SistemVaksinAPI/features/faskes"
	mocksFaskes "SistemVaksinAPI/features/faskes/mocks"
	"SistemVaksinAPI/features/vaksin"
	mocksVaksin "SistemVaksinAPI/features/vaksin/mocks"
	"os"
	"testing"
)

var (
	vaksinValue  vaksin.VaksinCore
	vaksinsValue []vaksin.VaksinCore
	vaksinData   mocksVaksin.Bussiness

	faskesData    mocksFaskes.Data
	faskesValue   faskes.FaskesCore
	faskessValue  []faskes.FaskesCore
	faskesUseCase faskes.Bussiness
)

func TestMain(m *testing.M) {
	faskesUseCase = NewFaskesBussiness(&faskesData, &vaksinData)

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

	os.Exit(m.Run())
}
