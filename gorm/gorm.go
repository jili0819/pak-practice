package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "root:mysqlps@123@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	infos := make([]TestTableA, 0)
	for i := 0; i < 1000000; i++ {
		v4uuid, err := uuid.NewV4()
		if err != nil {
			continue
		}
		infos = append(infos, TestTableA{Uuid: v4uuid.String()})
	}
	_ = db.CreateInBatches(&infos, 2000)
}

type TestTableA struct {
	Uuid       string    `gorm:"column:uuid;type:varchar(64);comment:列名字" json:"uuid"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`
}

func (m *TestTableA) TableName() string {
	return "test_table_a"
}
