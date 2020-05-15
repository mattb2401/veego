package veego

import "testing"

func TestVeego_RunServer(t *testing.T) {
	server := NewVeegoServer(".env", "env")
	if err := server.Run(); err != nil {
		t.Errorf("didn't expect any errors but got one: %v", err.Error())
	}
}
