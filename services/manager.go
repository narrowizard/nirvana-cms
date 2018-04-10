package services

import (
	"go-cms/meta"
	"go-cms/models"

	"github.com/caicloud/nirvana/log"
)

type ManagerService struct {
	BaseService
}

func NewManagerService() *ManagerService {
	return &ManagerService{
		BaseService: *NewBaseService(),
	}
}

func (this *ManagerService) UserList(page int, pagesize int, search string) ([]models.User, error) {
	var u = make([]models.User, 0)
	var err = this.DB.Model(&models.User{}).Where("account like ?", "%"+search+"%").Offset(pagesize * (page - 1)).Limit(pagesize).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	return u, nil
}

func (this *ManagerService) CreateUser(account, password string) error {
	var notFound = this.DB.Model(&models.User{}).Where("account=?", account).RecordNotFound()
	if !notFound {
		return meta.UserExistedError.Error(account)
	}
	var m models.User
	m.Account = account
	m.Password = password
	m.Encrypt()
	var err = this.DB.Model(&m).Create(&m).Error
	if err != nil {
		log.Error(err)
		return meta.TableInsertError.Error("users")
	}
	return nil
}
