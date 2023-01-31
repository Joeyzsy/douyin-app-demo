package initialize

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// Mysql 一个mysql库对于的结构体
type Mysql struct {
	Host      string
	Port      int32
	Database  string
	User      string
	Password  string
	Charset   string
	ParseTime bool `toml:"parse_time"`
}

type Redis struct {
	IP       string
	Port     int
	Database int
}

// Config 定义一个结构体接收toml所有配置
type Config struct {
	DB Mysql `toml:"mysql"`
	//RDB   Redis `toml:"redis"`
}

var Info Config

func init() {
	path, _ := os.Getwd()
	configPath := path + "/config/config_toml.toml"
	if _, err := toml.DecodeFile(configPath, &Info); err != nil {
		panic(err)
	}

	strings.Trim(Info.DB.Host, " ")
}

// DBConnectString 填充得到数据库连接字符串
func DBConnectString() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v",
		Info.DB.User, Info.DB.Password, Info.DB.Host, Info.DB.Port, Info.DB.Database,
		Info.DB.Charset, Info.DB.ParseTime)
	log.Println(arg)
	return arg
}
