package model

import "fmt"

type User struct {
	Id      float64
	Name    string
	Email   string
	Phone   string
	PwdHash string
	Pwd     string
	Active  bool
}

func (u *User) String() string {
	return fmt.Sprintf(`user{"Id": %v, "Name": "%v"}`, u.Id, u.Name)
}
