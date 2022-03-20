package main

import (
	"github.com/blademainer/go-exercise/demos/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	gorm.Init()
}
