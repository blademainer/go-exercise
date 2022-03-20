package gorm

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestNew(t *testing.T) {
	Init()
}
