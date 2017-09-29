package model

import ()

type Place struct {
	ID     int `gorm:primary_key`
	Name   string
	Town   Town
	TownId int `gorm:"ForeignKey:id"` //this tag didn't works
}


