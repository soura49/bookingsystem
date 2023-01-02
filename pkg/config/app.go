package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialets/mysql"
)

var (
	db *gorm.DB
)

// type dbConnectionString struct {
// 	dbUserName  string
// 	dbPassword  string
// 	dbTableName string
// }

// func getDbConnectionString(db *dbConnectionString) string {
// 	mySqlRequiredString := "?charset=utf8&parseTime=True&loc=Local"
// 	return db.dbUserName + ":" + db.dbPassword + "/" + db.dbTableName + mySqlRequiredString
// }

func Connect() {
	dbConnectionUrl := ""
	//dbConnectionUrl := getDbConnectionString(&dbConnectionString)
	d, err := gorm.Open("mysql", dbConnectionUrl)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
