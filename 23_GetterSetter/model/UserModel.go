package model

type UserModel struct {
	userName   string    
	email    string 
	password string 
}

func (user *UserModel) SetuserName(userName string) {
	user.userName = userName
}

func (user *UserModel) GetuserName() string {
	return user.userName
}

func (user *UserModel) Setemail(email string) {
	user.email = email
}

func (user UserModel) Getemail() string {
	return user.email
}
func (user *UserModel) Setpassword(password string) {
	user.password = password
}

func (user UserModel) GetPassword() string {
	return user.password
}
