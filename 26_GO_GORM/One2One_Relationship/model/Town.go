package model

import ()

type Town struct {
	ID   int `gorm:"primary_key"`
	Name string
}
