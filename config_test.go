package veego

import "testing"

func TestAppConfig_LoadEnv(t *testing.T) {
	config := NewAppConfig()
	conf, err := config.LoadEnv(".env")
	if err != nil {
		t.Errorf("didn't expect errorrs but got %v", err.Error())
	}
	if conf.Host != "0.0.0.0" {
		t.Errorf("TestAppConfig_LoadEnv expected 0.0.0.0 but got %v", conf.Host)
	}
}

func TestAppConfig_LoadYML(t *testing.T) {
	config := NewAppConfig()
	conf, err := config.LoadYML("config.yml")
	if err != nil {
		t.Errorf("didn't expect errorrs but got %v", err.Error())
	}
	if conf.Host != "0.0.0.0" {
		t.Errorf("TestAppConfig_LoadEnv expected 0.0.0.0 but got %v", conf.Host)
	}
}