package sql

import (
	"fmt"
	"github.com/Pauloricardo2019/teste_fazpay/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb(cfg *model.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.GetConnString()), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		fmt.Println("error opening connection ", err)
	}

	return db
}
