package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// GetConnection obtiene una conexión a la base de datos
func GetDB() (*gorm.DB, error) {
	var err error
	if db == nil {
		db, err = gorm.Open("mysql", "root:@/files?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Printf("########## DB ERROR: " + err.Error() + " #############")
			db = nil
		}
	}

	return db, err
}
