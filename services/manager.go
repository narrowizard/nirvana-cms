package services

import (
	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"

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
	var u = make([]models.UserInfo, 0)
	err = this.DB.Table("users u").Joins("join roles r on u.role_id = r.id").Select("u.id,role_id,account,u.created_at,u.status,u.updated_at,r.name as role_name,r.status as role_status").Where("u.status < 100 and account like ?", "%"+search+"%").Offset(pagesize * (page - 1)).Limit(pagesize).Scan(&u).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	return &models.PagableData{
		Total: count,
		Data:  u,
	}, nil
}

func (this *ManagerService) CreateUser(account, password string, roleid int) (*models.User, error) {
	var m models.User
	var notFound = this.DB.Model(&models.User{}).Where("account=?", account).Scan(&m).RecordNotFound()
	if !notFound {
		return nil, meta.UserExistedError.Error(account)
	}
	m.Account = account
	m.Password = password
	m.RoleID = roleid
	m.Status = models.STATUSNORMAL
	m.Encrypt()
	var err = this.DB.Model(&m).Create(&m).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableInsertError.Error("users")
	}
	m.ClearPassword()
	return &m, nil
}

func (this *ManagerService) UpdateUser(userid int, status models.ENUMSTATUS, roleid int) error {
	var fields = make(map[string]interface{})
	var needUpdate = false
	if roleid > 0 {
		needUpdate = true
		fields["role_id"] = roleid
	}
	if status > 0 {
		needUpdate = true
		fields["status"] = int(status)
	}
	if !needUpdate {
		return nil
	}
	var err = this.DB.Model(&models.User{}).Where("id=?", userid).Update(fields).Error
	if err != nil {
		log.Error(err)
		return meta.TableUpdateError.Error("users")
	}
	return nil
}

func (this *ManagerService) MenuList() ([]meta.SimpleTreeNode, error) {
	var menuData = make([]models.Menu, 0)
	var err = this.DB.Model(&models.Menu{}).Where("status < 100").Scan(&menuData).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("menus")
	}
	// 转换成树形结构
	var treeableData = make([]meta.TreeableData, 0)
	for _, v := range menuData {
		treeableData = append(treeableData, v)
	}
	var treeData = meta.ArrayToTree(treeableData)
	return treeData, nil
}

func (this *ManagerService) CreateMenu(pid, ismenu int, name, url, icon, remarks string) error {
	var m = models.Menu{
		Icon:     icon,
		Name:     name,
		URL:      url,
		ParentID: pid,
		Status:   models.STATUSNORMAL,
		IsMenu:   ismenu,
		Remarks:  remarks,
	}
	var err = this.DB.Create(&m).Error
	if err != nil {
		log.Error(err)
		return meta.TableInsertError.Error("menus")
	}
	return nil
}

func (this *ManagerService) UpdateMenu(id, ismenu int, name, url, icon, remarks string) error {
	var err = this.DB.Model(&models.Menu{}).Where("id=? and status=?", id, models.STATUSNORMAL).Update(map[string]interface{}{
		"name":    name,
		"url":     url,
		"icon":    icon,
		"is_menu": ismenu,
		"remarks": remarks,
	}).Error
	if err != nil {
		log.Error(err)
		return meta.TableUpdateError.Error("menus")
	}
	return nil
}

func (this *ManagerService) UserInfo(id int) (*models.UserInfo, error) {
	var m models.UserInfo
	var err = this.DB.Table("users u").Joins("join roles r on u.role_id = r.id").Where("u.id=? and u.status<?", id, models.STATUSDELETED).Select("u.id,account,role_id,r.name as role_name,r.status as role_status,u.status").Scan(&m).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	return &m, nil
}
