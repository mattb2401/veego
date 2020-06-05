package veego

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type databaseManager struct {
	databaseURL string
}

type dBParams struct {
	Schema   string
	Host     string
	Username string
	Password string
	Database string
	Port     string
}

func NewDatabaseManager(databaseURL string) *databaseManager {
	return &databaseManager{
		databaseURL: databaseURL,
	}
}

func (d *databaseManager) Connect() (*gorm.DB, error) {
	params, err := d.urlParser()
	if err != nil {
		return nil, err
	}
	switch params.Schema {
	case "mysql":
		db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true`, params.Username, params.Password, params.Host, params.Port, params.Database))
		if err != nil {
			return nil, err
		}
		return db, nil
	case "postgres":
		db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", params.Host, params.Port, params.Username, params.Database, params.Password))
		if err != nil {
			return nil, err
		}
		return db, nil
	case "mssql":
		db, err := gorm.Open("mssql", fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", params.Username, params.Password, params.Host, params.Port, params.Database))
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, errors.New("unknown Database schema")
	}
}

func (d *databaseManager) urlParser() (*dBParams, error) {
	var host, port, path, password string
	var err error
	u, err := url.Parse(d.databaseURL)
	if err != nil {
		return &dBParams{}, err
	}
	if host, port, err = net.SplitHostPort(u.Host); err == nil {
	} else {
		host = u.Host
	}
	if strings.Contains(u.Path, "/") {
		path = strings.Split(u.Path, "/")[1]
	} else {
		path = u.Path
	}
	if pwd, ok := u.User.Password(); ok {
		password = pwd
	}
	return &dBParams{
		Schema:   u.Scheme,
		Username: u.User.Username(),
		Password: password,
		Host:     host,
		Database: path,
		Port:     port,
	}, nil
}
