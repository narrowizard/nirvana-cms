package services

import (
	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"

	"github.com/caicloud/nirvana/log"
)

type UserService struct {
	BaseService
}

func NewUserService() *UserService {
	return &UserService{
		BaseService: *NewBaseService(),
	}
}

func (this *UserService) MenuList(userid int) ([]meta.SimpleTreeNode, error) {
	var menuData = make([]models.Menu, 0)
	var err = this.DB.Raw(`select m.* from menus m 
									join role_menus rm on m.id = rm.menu_id
									join users u on rm.role_id = u.role_id
									join roles r on rm.role_id = r.id
									where u.id=? and rm.status=? and m.status=? and r.status=? and m.is_menu=1`, userid, models.STATUSNORMAL, models.STATUSNORMAL, models.STATUSNORMAL).Scan(&menuData).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("user_menus")
	}
	// 转换成树形结构
	var treeableData = make([]meta.TreeableData, 0)
	for _, v := range menuData {
		treeableData = append(treeableData, v)
	}
	var treeData = meta.ArrayToTree(treeableData)
	return treeData, nil
}
