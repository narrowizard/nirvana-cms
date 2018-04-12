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

func (this *ManagerService) UserList(page int, pagesize int, search string) (*models.PagableData, error) {
	var count int
	var err = this.DB.Model(&models.User{}).Where("status < 100 and account like ?", "%"+search+"%").Count(&count).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	var u = make([]models.User, 0)
	err = this.DB.Model(&models.User{}).Select("id,account,created_at,status,updated_at").Where("status < 100 and account like ?", "%"+search+"%").Offset(pagesize * (page - 1)).Limit(pagesize).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	return &models.PagableData{
		Total: count,
		Data:  u,
	}, nil
}

func (this *ManagerService) CreateUser(account, password string) error {
	var m models.User
	var notFound = this.DB.Model(&models.User{}).Where("account=?", account).Scan(&m).RecordNotFound()
	if !notFound {
		return meta.UserExistedError.Error(account)
	}
	m.Account = account
	m.Password = password
	m.Status = int(models.STATUSNORMAL)
	m.Encrypt()
	var err = this.DB.Model(&m).Create(&m).Error
	if err != nil {
		log.Error(err)
		return meta.TableInsertError.Error("users")
	}
	return nil
}

func (this *ManagerService) UpdateUser(userid int, status models.ENUMSTATUS) error {
	var err = this.DB.Model(&models.User{}).Where("id=?", userid).Update(map[string]interface{}{
		"status": int(status),
	}).Error
	if err != nil {
		log.Error(err)
		return meta.TableUpdateError.Error("users")
	}
	return nil
}
