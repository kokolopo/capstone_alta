package client

type IService interface {
	AddClient(userID int, input InputAddClient) (Client, error)
	GetAll(userID int) ([]Client, error)
}

type service struct {
	repository IRepository
}

func NewUserService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) AddClient(userID int, input InputAddClient) (Client, error) {
	var data Client

	data.Fullname = input.Fullname
	data.Email = input.Email
	data.Address = input.Address
	data.City = input.City
	data.ZipCode = input.ZipCode
	data.Company = input.Company
	data.UserID = userID

	client, err := s.repository.Save(data)
	if err != nil {
		return client, err
	}

	return client, nil
}

func (s *service) GetAll(userID int) ([]Client, error) {
	clients, err := s.repository.FindAll(userID)
	if err != nil {
		return clients, err
	}

	return clients, nil
}
