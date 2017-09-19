package config

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DbConfig ...
type DbConfig struct {
	DBMS     string `json:"dbms"`
	PROTOCOL string `json:"protocol"`
	HOST     string `json:"host"`
	PORT     int    `json:"port"`
	USER     string `json:"user"`
	PASS     string `json:"pass"`
	DBNAME   string `json:"db_name"`
}

// NewDbConfig ...
func NewDbConfig(j string) *DbConfig {
	file, err := ioutil.ReadFile(j)
	if err != nil {
		panic(err)
	}
	config := DbConfig{}
	json.Unmarshal(file, &config)
	return &config
}

// GormOpen ...
func (config *DbConfig) GormOpen() (db *gorm.DB, err error) {

	// dbms , user:pass@protocol(host:port)/dbname
	db, err = gorm.Open(config.DBMS, (config.USER + ":" + config.PASS + "@" + config.PROTOCOL + "(" + config.HOST + ":" + strconv.Itoa(config.PORT) + ")" + "/" + config.DBNAME + "?charset=utf8&parseTime=True&loc=Local"))
	db.LogMode(true)

	return
}
