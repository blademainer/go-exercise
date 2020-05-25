package gorm

import "testing"
import _ "github.com/jinzhu/gorm/dialects/sqlite"


func TestNew(t *testing.T) {
	Init()
}