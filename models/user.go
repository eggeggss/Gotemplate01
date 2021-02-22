package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// User is a struct
type User struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var (
	Users     []User
	SqlString *string
)

func init() {

	Users = []User{
		{ID: 1, Name: "Duotify"},
		{ID: 2, Name: "Duotify"},
	}
	readConfig()

}

type Conf struct {
	ConnectionString string `json:"ConnectionString"`
}

func readConfig() {

	file, _ := os.Open("./conf.json")

	content, _ := ioutil.ReadAll(file)

	var u Conf
	err := json.Unmarshal(content, &u)
	if err != nil {
		log.Print(err)
	}
	SqlString = &u.ConnectionString

}
