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

func (this *ManagerService) CreateUser(account, password string, menus []int) (*models.User, error) {
	var m models.User
	var notFound = this.DB.Model(&models.User{}).Where("account=?", account).Scan(&m).RecordNotFound()
	if !notFound {
		return nil, meta.UserExistedError.Error(account)
	}
	m.Account = account
	m.Password = password
	m.Status = models.STATUSNORMAL
	m.Encrypt()
	var tx = this.DB.Begin()
	var err = tx.Model(&m).Create(&m).Error
	if err != nil {
		log.Error(err)
		tx.Rollback()
		return nil, meta.TableInsertError.Error("users")
	}
	for _, v := range menus {
		var um = models.UserMenu{
			UserID: m.ID,
			MenuID: v,
			Status: models.STATUSNORMAL,
		}
		err = tx.Model(&models.UserMenu{}).Create(&um).Error
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, meta.TableInsertError.Error("user_menus")
		}
	}
	tx.Commit()
	m.ClearPassword()
	return &m, nil
}

func (this *ManagerService) UpdateUser(userid int, status models.ENUMSTATUS, menus []int) error {
	var err error
	var tx = this.DB.Begin()
	if status > 0 {
		err = tx.Model(&models.User{}).Where("id=?", userid).Update(map[string]interface{}{
			"status": int(status),
		}).Error
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return meta.TableUpdateError.Error("users")
		}
	}
	if menus != nil {
		// 首先清空原本的菜单
		err = tx.Model(&models.UserMenu{}).Where("user_id=? and status=?", userid, models.STATUSNORMAL).Update(map[string]interface{}{
			"status": models.STATUSDELETED,
		}).Error
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return meta.TableUpdateError.Error("user_menus")
		}
		for _, v := range menus {
			// 创建新菜单
			var temp = models.UserMenu{
				UserID: userid,
				Status: models.STATUSNORMAL,
				MenuID: v,
			}
			err = tx.Model(&models.UserMenu{}).Create(&temp).Error
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return meta.TableInsertError.Error("user_menus")
			}
		}
	}
	tx.Commit()
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
	var err = this.DB.Model(&models.User{}).Where("id=? and status<?", id, models.STATUSDELETED).Select("id,account").Scan(&m).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("users")
	}
	err = this.DB.Model(&models.UserMenu{}).Where("user_id=? and status=?", id, models.STATUSNORMAL).Scan(&m.Menus).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("user_menus")
	}
	return &m, nil
}
