package models

import (
	"errors"
	"gorm.io/gorm"
	"sample-api-rest/storage"
)

var Customer *customer

type customer struct {
	CustomerId       string
	CustomerIdNumber string
	CustomerName     string
}

func (t *customer) Create() ([]customer, error) {
	var data = []customer{}

	db := storage.GetDBInstance()

	if err := db.Table("customer").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (t *customer) FindAll() ([]customer, error) {
	var data = []customer{}

	db := storage.GetDBInstance()

	if err := db.Table("customer").Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (t *customer) FindById(id string) (customer, error) {
	var data customer

	db := storage.GetDBInstance()

	if err := db.Table("customer").Where("customer_id = ?", id).Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer{}, errors.New("customer does not exist")
		}
		return customer{}, errors.New("internal error")
	}

	return data, nil
}
