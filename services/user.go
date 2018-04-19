package services

import (
	"nirvana-cms/meta"
	"nirvana-cms/models"

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
									join user_menus um on m.id = um.menu_id
									where um.user_id=? and um.status=? and m.status=? and m.is_menu=1`, userid, models.STATUSNORMAL, models.STATUSNORMAL).Scan(&menuData).Error
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
