package model

type Config struct {
	Environment  string
	DbConnString string
	RestPort     int
	EnableSentry bool
	BasePath     string
}
