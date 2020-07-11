package models

type Auth struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, passowrd string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: passowrd}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
