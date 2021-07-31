package gorm

import "github.com/blademainer/commons/pkg/logger"

type Person struct {
	Name string `json:"name" gorm:"column:name,index:name_idx"`
	Age  int    `json:"age"`
}

func (p *Person) TableName() string {
	return "person"
}

func Hook() {
	db, err := InitDb()
	if err != nil {
		logger.Fatal(err)
	}
	err = db.CreateTable(&Person{}).Error
	if err != nil {
		logger.Fatal(err.Error())
	}

}
