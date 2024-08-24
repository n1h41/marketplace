package config

import "testing"

func TestGetConfig(t *testing.T) {
	config := GetConfig()
	t.Log(config.Server.Port)
	t.Log(config.Db.DbName)
}
