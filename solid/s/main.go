package main

import "fmt"

func main() {

	u := Admin{}
	u.setLogin("Admin")

	isAuth := Auth(u, UserFinder{})

	fmt.Println(isAuth)

}

// UserFinder ...
type UserFinder struct{}

func (uf *UserFinder) find(u Admin) bool {
	loginInStorage := "Admin"
	return loginInStorage == u.getLogin()
}

// Admin ...
type Admin struct {
	name, lastName, login, email string
}

func (u *Admin) getLogin() string {
	return u.login
}

func (u *Admin) setLogin(l string) {
	u.login = l
}

// Auth ...
func Auth(u Admin, f UserFinder) bool {
	return f.find(u)
}
