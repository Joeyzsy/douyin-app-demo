package initialize

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/RaymondCode/simple-demo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func initDB() {
	fmt.Printf("Start initDB")
	var err error
	//gorm.Open(mysql.Open(DBConnectString()),&gorm.Config)

	global.DB, err = gorm.Open(mysql.Open(DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数", err)
	}

}
