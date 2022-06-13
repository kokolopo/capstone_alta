package database

import (
	"github.com/kokolopo/capstone_alta/domain/client"
	"github.com/kokolopo/capstone_alta/domain/user"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: user.User{}},
		{Model: client.Client{}},
	}
}
