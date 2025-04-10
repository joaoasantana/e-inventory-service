package config

import "fmt"

type DatabaseInfo struct {
	Driver string
	Host   string
	Port   string
	Name   string
	User   string
	Pass   string
}

func (d *DatabaseInfo) URL() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Pass, d.Name,
	)
}
