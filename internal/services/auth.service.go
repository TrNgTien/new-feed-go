package services

type AuthService interface {
	Login(username, password string) (map[string]interface{}, error)
	Register(username, password string) (bool, error)
	Logout() (bool, error)
}
type authServiceImpl struct{}

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

func (as *authServiceImpl) Login(username, password string) (map[string]interface{}, error) {
	userData := make(map[string]interface{})
	userData["username"] = username
	userData["password"] = password

	return userData, nil
}

func (as *authServiceImpl) Register(username, password string) (bool, error) {
	return true, nil
}

func (as *authServiceImpl) Logout() (bool, error) {
	return true, nil
}
