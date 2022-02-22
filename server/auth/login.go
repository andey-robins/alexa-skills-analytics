package auth

type LoginService interface {
	LoginWebUser(email, password string) bool
}

type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "test@example.com",
		password: "password",
	}
}

func (info *loginInformation) LoginWebUser(email, password string) bool {
	return info.email == email && info.password == password
}
