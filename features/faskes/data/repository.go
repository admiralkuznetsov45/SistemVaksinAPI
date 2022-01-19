package data

import (
	"SistemVaksinAPI/features/faskes"

	"gorm.io/gorm"
)

type mysqlFaskesRepository struct {
	Conn *gorm.DB
}

func NewFaskesRepository(conn *gorm.DB) faskes.Data {
	return &mysqlFaskesRepository{
		Conn: conn,
	}
}

func (dr *mysqlFaskesRepository) InsertFaskes(data faskes.FaskesCore) (resp faskes.FaskesCore, err error) {

	record := fromCore(data)
	if err := dr.Conn.Create(&record).Error; err != nil {

		return faskes.FaskesCore{}, err
	}

	return toCore(&record), nil
}

func (dr *mysqlFaskesRepository) SelectAllFaskes() (resp []faskes.FaskesCore) {

	var record []Faskes

	if err := dr.Conn.Model(&Faskes{}).Find(&record).Error; err != nil {
		return []faskes.FaskesCore{}
	}

	return toList(record)
}

func (dr *mysqlFaskesRepository) SelectFaskesByID(ID int) (resp faskes.FaskesCore, err error) {
	var record Faskes

	if err := dr.Conn.First(&record, ID).Error; err != nil {
		return faskes.FaskesCore{}, err
	}

	return toCore(&record), nil
}

// func (dr *mysqlFaskesRepository) SelectFaskesByName(Nama string) (resp faskes.FaskesCore, err error) {
// 	var record Faskes

// 	if err := dr.Conn.First(&record, Nama).Error; err != nil {
// 		return faskes.FaskesCore{}, err
// 	}

// 	return toCore(&record), nil
// }

func (dr *mysqlFaskesRepository) SelectFaskesByName(data faskes.FaskesCore) (resp faskes.FaskesCore, err error) {

	record := fromCore(data)

	if err := dr.Conn.Model(&Faskes{}).Where("nama = ?", data.Nama).First(&record).Error; err != nil {
		return faskes.FaskesCore{}, err
	}

	if err != nil {
		return faskes.FaskesCore{}, err
	}

	if err := dr.Conn.Model(&Faskes{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return faskes.FaskesCore{}, err
	}

	// fmt.Println(record.Token)
	return toCore(&record), err
}
