package auth

import "fmt"

type User struct {
	Name  string
	Phone int
}

func (u User) GetData() []string {
	return []string{u.Name, fmt.Sprintf("%d", u.Phone)}
}
