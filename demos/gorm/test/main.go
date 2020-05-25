package main

import _ "github.com/jinzhu/gorm/dialects/sqlite"
import _ "github.com/jinzhu/gorm/dialects/mysql"
import "github.com/blademainer/go-exercise/demos/gorm"

func main() {
	gorm.Init()
}
