package client

import "gorm.io/gorm"

type IRepository interface {
	Save(client Client) (Client, error)
	FindAll(userID int) ([]Client, error)
	FindById(id int) (Client, error)
	FindByEmail(email string) (Client, error)
	Update(client Client) (Client, error)
	Delete(client Client) (Client, error)
}

type repository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(client Client) (Client, error) {
	err := r.DB.Create(&client).Error
	if err != nil {
		return client, err
	}

	return client, nil
}

func (r *repository) FindAll(userID int) ([]Client, error) {
	var clients []Client
	err := r.DB.Where("user_id = ?", userID).Find(&clients).Error
	if err != nil {
		return clients, err
	}

	return clients, nil
}

func (r *repository) FindById(id int) (Client, error) {
	var client Client
	err := r.DB.Where("id = ?", id).Find(&client).Error
	if err != nil {
		return client, err
	}

	return client, nil
}

func (r *repository) FindByEmail(email string) (Client, error) {
	var client Client

	err := r.DB.Where("email = ?", email).Find(&client).Error
	if err != nil {
		return client, err
	}

	return client, nil
}

func (r *repository) Update(client Client) (Client, error) {
	err := r.DB.Save(&client).Error

	if err != nil {
		return client, err
	}

	return client, nil
}

func (r *repository) Delete(client Client) (Client, error) {
	err := r.DB.Delete(&client).Error
	if err != nil {
		return client, err
	}

	return client, nil
}
