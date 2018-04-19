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

func (this *UserService) MenuList(userid int) ([]models.Menu, error) {
	var data = make([]models.Menu, 0)
	var err = this.DB.Raw(`select m.* from menus m 
									join user_menus um on m.id = um.menu_id
									where um.user_id=? and um.status=? and m.status=? and m.is_menu=1`, userid, models.STATUSNORMAL, models.STATUSNORMAL).Scan(&data).Error
	if err != nil {
		log.Error(err)
		return nil, meta.TableQueryError.Error("user_menus")
	}
	return data, nil
}
