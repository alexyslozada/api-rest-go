package usuario

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var storage Storage

func init() {
	storage = make(map[string]*Model)
	u1 := &Model{
		FirstName: "Alexys",
		Email:     "alexys@ed.team",
		Password:  "123456",
	}
	u2 := &Model{
		FirstName: "Juan",
		Email:     "juan@ed.team",
		Password:  "123456",
	}
	storage.Create(u1)
	storage.Create(u2)
}

type Model struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Storage map[string]*Model

func (s Storage) Create(m *Model) *Model {
	s[m.Email] = m
	return s[m.Email]
}

func (s Storage) GetAll() Storage {
	return s
}

func (s Storage) GetAllPaginate(l, p int) []*Model {
	us := make([]*Model, 0, len(s))
	for _, v := range s {
		us = append(us, v)
		fmt.Println(v)
	}
	fmt.Println(us)
	offset := l*p - l
	r := us[offset : l*p]
	return r
}

func (s Storage) GetByEmail(e string) *Model {
	if v, ok := s[e]; ok {
		return v
	}

	return nil
}

func (s Storage) Delete(e string) {
	delete(s, e)
}

func (s Storage) Update(e string, z *Model) *Model {
	s[e] = z
	return s[e]
}

func (s Storage) Login(e, p string) *Model {
	for _, v := range s {
		if v.Email == e && v.Password == p {
			return v
		}
	}

	return nil
}

type Claim struct {
	Usuario Model
	jwt.StandardClaims
}
