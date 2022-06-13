package user

import (
	"errors"
	"log"

	utl "github.com/kokolopo/capstone_alta/utils"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Register(input InputRegister) (User, error)
	Login(input InputLogin) (User, error)
	GetUserById(id int) (User, error)
	IsEmailAvailable(input InputCheckEmail) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	UpdateUser(id int, input InputUpdate) (User, error)
	ResetPassword(input InputCheckEmail) (User, error)
}

type service struct {
	repository IRepository
}

func NewUserService(repository IRepository) *service {
	return &service{repository}
}

func (s *service) Register(input InputRegister) (User, error) {
	var newUser User

	// cek email
	user, err := s.repository.FindByEmail(input.Email)
	if user.ID != 0 {
		return user, err
	}

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return newUser, errHash
	}

	//tangkap nilai dari inputan
	newUser.Fullname = input.Fullname
	newUser.Email = input.Email
	newUser.NoTlpn = input.NoTlpn
	newUser.BusinessName = input.BusinessName
	newUser.Password = string(passwordHash)
	newUser.Role = "user"
	newUser.Avatar = "images/default_user.png"

	//save data yang sudah dimapping kedalam struct Mahasiswa
	user, err = s.repository.Save(newUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input InputCheckEmail) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) Login(input InputLogin) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	//cek jika user tidak ada
	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	return user, nil
}

func (s *service) SaveAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdateUser(id int, input InputUpdate) (User, error) {
	user, errUser := s.repository.FindById(id)
	if errUser != nil {
		return user, errUser
	}

	user.ID = id
	user.Fullname = input.Fullname
	user.Email = input.Email
	user.NoTlpn = input.NoTlpn
	user.BusinessName = input.BusinessName

	updatedUser, errUpdate := s.repository.Update(user)
	if errUpdate != nil {
		return updatedUser, errUpdate
	}

	return updatedUser, nil
}

func (s *service) ResetPassword(input InputCheckEmail) (User, error) {
	email := input.Email

	// cek apakah terdapat data user dengan email tersebut
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, err
	}

	// generate password
	res, err := password.Generate(10, 7, 3, false, false)
	if err != nil {
		log.Fatal(err)
	}

	//enkripsi password
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(res), bcrypt.MinCost)

	user.Password = string(passwordHash)

	updatePass, errPass := s.repository.Update(user)
	if errPass != nil {
		return user, err
	} else {
		utl.SendMail(email, res)
	}

	return updatePass, nil
}
