package dal

import (
	"encoding/json"
	"fmt"
	"hw/models"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	dsn *string
)

func init() {
	file, _ := os.Open("./conf.json")

	content, _ := ioutil.ReadAll(file)

	var u models.Conf
	err := json.Unmarshal(content, &u)
	if err != nil {
		log.Print(err)
	}
	dsn = &u.ConnectionString
}

func Gets() []models.UserTable {

	db, err := gorm.Open(sqlserver.Open(*dsn), &gorm.Config{})

	fmt.Println(db, err)
	return nil
}
