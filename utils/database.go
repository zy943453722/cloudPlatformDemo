package utils

import (
	"cloudPlatformDemo/middleware"
	"fmt"
	"github.com/go-xorm/xorm"
)

var Db *xorm.Engine

func xormInit(mysqlConfig map[string]string) error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		mysqlConfig["username"],
		mysqlConfig["password"],
		mysqlConfig["host"] + ":" + mysqlConfig["port"],
		mysqlConfig["database"],
		mysqlConfig["charset"],
	)
	engine, err := xorm.NewEngine(mysqlConfig["driver"], connStr)
	if err != nil {
		return err
	}
	logger := xorm.NewSimpleLogger(middleware.Log.Out)
	logger.ShowSQL(true)
	engine.SetLogger(logger)
	
	Db = engine
	return nil
}

func GetDb(mysqlConfig map[string]string) (*xorm.Engine, error) {
	if err := xormInit(mysqlConfig); err != nil {
		return nil, err
	}

	return Db, nil
}
