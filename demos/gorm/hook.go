package gorm

import (
	"fmt"
	"github.com/blademainer/go-exercise/pkg/time"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/sqlite"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Person struct {
	Name string `json:"name" gorm:"column:name;primary_key"`
	Age  int    `json:"age" gorm:"column:age"`
}

func (u *Person) BeforeSave() (err error) {
	fmt.Printf("before creating: %v\n", u)
	return
}

func (u *Person) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Printf("after create person: %v scope: %v \n", u, scope)
	return
}
func init() {
	gorm.DefaultCallback.Create().Before("gorm:before_create").Register("test:create", func(scope *gorm.Scope) {
		fmt.Printf("callback create, scope.Value: %v, scope.Search: %v, scope.SQL: %v, scope.SQLVars: %v \n", scope.Value, scope.Search, scope.SQL, scope.SQLVars)
	})
	gorm.DefaultCallback.Update().Register("test:update", func(scope *gorm.Scope) {
		fmt.Printf("callback update, scope.Value: %v, scope.Search: %v, scope.SQL: %v, scope.SQLVars: %v \n", scope.Value, scope.Search, scope.SQL, scope.SQLVars)
	})
	gorm.DefaultCallback.Query().Register("test:query", func(scope *gorm.Scope) {
		fmt.Printf("callback query, scope.Value: %v, scope.Search: %v, scope.SQL: %v, scope.SQLVars: %v \n", scope.Value, scope.Search, scope.SQL, scope.SQLVars)
	})
	gorm.DefaultCallback.Delete().Register("test:delete", func(scope *gorm.Scope) {
		fmt.Printf("callback delete, scope.Value: %v, scope.Search: %v, scope.SQL: %v, scope.SQLVars: %v \n", scope.Value, scope.Search, scope.SQL, scope.SQLVars)
	})
}

func Init() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	err = db.CreateTable(&Person{}).Error
	if err != nil {
		//panic(err)
	}

	err = db.Create(&Person{Name: time.NowTimeString()}).Error
	if err != nil {
		panic(err)
	}

	people := make([]*Person, 0)
	err = db.Model(&Person{}).Scan(&people).Error
	if err != nil {
		panic(err)
	}

}
