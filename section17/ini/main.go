package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

// ini configファイルを読み込むのに使用される

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(8080), // Mustで指定した場合config.iniにない場合読み込まれる

		DbName: cfg.Section("db").Key("name").MustString("example.com"),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
