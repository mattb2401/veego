package veego

import (
	"testing"
)

func TestDatabase_URLParser(t *testing.T) {
	config := NewAppConfig()
	conf, err := config.LoadEnv(".env")
	dbManager := NewDatabaseManager(conf.DatabaseURL)
	params, err := dbManager.urlParser()
	if err != nil {
		t.Errorf("didn't expect any errors but here we are : %v", err.Error())
	}
	if params.Schema != "mysql" {
		t.Errorf("expected schema mysql but got something else: %v", params.Schema)
	}
}

func TestDatabase_Connect(t *testing.T) {
	config := NewAppConfig()
	conf, err := config.LoadEnv(".env")
	dbManager := NewDatabaseManager(conf.DatabaseURL)
	_, err = dbManager.Connect()
	if err != nil {
		t.Errorf("didn't expect any errors but here we are : %v", err.Error())
	}
}
