package conf

import (
	"testing"
)

//
func TestLoadConfig(t *testing.T) {
	file := "./v1.json"
	c := GetConfig(file)
	if c == nil {
		t.Error("Get nil config")
	}
	if c.Mysql.Port == 0 {
		t.Error("Database port is 0")
	}
	if (c.Mysql.GetConnStr()) == "" {
		t.Error("mysql connstr is nil ")
	}
}
