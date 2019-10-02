package controller

import (
	"github.com/go-pg/pg/v9"
	"yummy/config"
)

type Controller struct {
	DB *pg.DB
	Config config.Config
}

func NewController() *Controller {
	var c Controller
	return &c 
}