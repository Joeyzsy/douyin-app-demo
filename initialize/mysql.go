package initialize

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
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

// DBConnectString 填充得到数据库连接字符串
func DBConnectString() string {
	username := "root"
	password := "123456"
	host := "localhost"
	port := 3306
	dbname := "douyin"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)

	return dsn
}

/*初始化数据库*/
func InitDB() {

	fmt.Printf(DBConnectString())
	var err error
	//gorm.Open(mysql.Open(initialize.DBConnectString()), &gorm.Config)

	global.DB, err = gorm.Open(mysql.Open(DBConnectString()), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
		//Logger:                 logger.Default.LogMode(logger.Info), //打印sql语句
	})
	if err != nil {
		fmt.Print("error!")
	}
	//测试连接
	/*
		var users []model.Users
		result := global.DB.Find(&users)
		fmt.Println("********************")
		fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
		fmt.Printf("result.RowsAffected: %v\n", users[0])
		fmt.Println("********************")
	*/
	var videos []model.Video
	result := global.DB.Find(&videos)
	fmt.Println("********************")
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
	fmt.Printf("result.RowsAffected: %v\n", videos[0])
	fmt.Println("********************")
}
