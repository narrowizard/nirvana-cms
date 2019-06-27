package services

import (
	"github.com/caicloud/nirvana/log"
	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"
)

type RoleService struct {
	BaseService
}

func NewRoleService() *RoleService {
	return &RoleService{
		BaseService: *NewBaseService(),
	}
}

func (r *RoleService) List(search string) ([]models.Role, error) {
	var m = make([]models.Role, 0)
	var err = r.DB.Model(&models.Role{}).Where("status < ?", models.STATUSDELETED).Where("name like ?", "%"+search+"%").Scan(&m).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("roles")
	}
	return m, nil
}

func (r *RoleService) Info(id int) (*models.RoleInfo, error) {
	var m models.RoleInfo
	var err = r.DB.Model(&models.Role{}).Where("id=?", id).Scan(&m).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("roles")
	}
	m.Menus = make([]models.RoleMenu, 0)
	err = r.DB.Model(&models.RoleMenu{}).Where("role_id=? and status=?", id, models.STATUSNORMAL).Scan(&m.Menus).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("role_menus")
	}
	return &m, nil
}

func (r *RoleService) Update(id int, name string, menus []int, status int) error {
	var role models.Role
	var db = r.DB.Model(&models.Role{}).Where("name=? and id <> ?", name, id).Scan(&role)
	var notFound = db.RecordNotFound()
	var err = db.Error
	if notFound {
		var tx = r.DB.Begin()
		var fields = make(map[string]interface{}, 0)
		var needUpdate = false
		if name != "" {
			needUpdate = true
			fields["name"] = name
		}
		if status > 0 {
			needUpdate = true
			fields["status"] = status
		}
		if needUpdate {
			err = tx.Model(&models.Role{}).Where("id=?", id).Update(fields).Error
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return meta.TableUpdateError.Error("roles")
			}
		}
		if menus != nil {
			err = tx.Model(&models.RoleMenu{}).Where("role_id=? and status=?", id, models.STATUSNORMAL).Update(map[string]interface{}{
				"status": models.STATUSDELETED,
			}).Error
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return meta.TableUpdateError.Error("role_menus")
			}
			for _, v := range menus {
				var temp = &models.RoleMenu{
					RoleID: id,
					MenuID: v,
					Status: models.STATUSNORMAL,
				}
				err = tx.Model(&temp).Create(&temp).Error
				if err != nil {
					log.Error(err)
					tx.Rollback()
					return meta.TableInsertError.Error("role_menus")
				}
			}
		}
		tx.Commit()
		return nil
	}
	if err != nil {
		log.Error(err)
		return meta.TableQueryError.Error("roles")
	} else {
		return meta.RoleExistedError.Error(name)
	}
}

func (r *RoleService) New(name string, menus []int) (*models.Role, error) {
	var role models.Role
	var db = r.DB.Model(&models.Role{}).Where("name=?", name).Scan(&role)
	var notFound = db.RecordNotFound()
	var err = db.Error
	if notFound {
		var m models.Role
		m.Name = name
		m.Status = models.STATUSNORMAL
		var tx = r.DB.Begin()
		err = tx.Model(&m).Create(&m).Error
		if err != nil {
			log.Error(err)
			tx.Rollback()
			return nil, meta.TableInsertError.Error("roles")
		}
		for _, v := range menus {
			var temp = &models.RoleMenu{
				RoleID: m.ID,
				MenuID: v,
				Status: models.STATUSNORMAL,
			}
			err = tx.Model(&temp).Create(&temp).Error
			if err != nil {
				log.Error(err)
				tx.Rollback()
				return nil, meta.TableInsertError.Error("role_menus")
			}
		}
		tx.Commit()
		return &m, nil
	}
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("roles")
	} else {
		return nil, meta.RoleExistedError.Error(name)
	}
}
