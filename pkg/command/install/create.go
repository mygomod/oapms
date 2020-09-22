package install

import (
	"oapms/pkg/model"
	"oapms/pkg/mus"
)

func Create() error {
	db := mus.Db
	models := []interface{}{
		&model.User{},
		&model.UserSecret{},
		&model.Access{},
		&model.App{},
		&model.Authorize{},
		&model.CasbinRule{},
		&model.Department{},
		&model.Expires{},
		&model.Menu{},
		&model.MenuPms{},
		&model.Pms{},
		&model.Refresh{},
		&model.Role{},
		&model.RolePms{},
	}
	db.SingularTable(true)
	if db.Error != nil {
		return db.Error
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
