package app

import (
	"log"
)

type Context struct {
	dbGateway   *DbGateway
	tplRenderer *TplRenderer
	*log.Logger
}

func NewContext(dbGateway *DbGateway, renderer *TplRenderer, logger *log.Logger) Context {
	return Context{
		dbGateway,
		renderer,
		logger,
	}
}
