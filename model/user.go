package model

type User struct {
	ID           int    `json:"id" xorm:"pk autoincr"`
	Name         string `json:"name"`
	PasswordHash []byte `json:"-"`
}
