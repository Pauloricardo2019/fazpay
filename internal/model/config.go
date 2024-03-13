package model

import "fmt"

type Config struct {
	Environment  string
	DBConfig     DBConfig
	RestPort     int
	EnableSentry bool
	BasePath     string
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (d *DBConfig) GetConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Name)
}
