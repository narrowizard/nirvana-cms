package services

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	return &UserService{
		BaseService: *NewBaseService(),
	}
}

func (this *UserService) Login() {

}
