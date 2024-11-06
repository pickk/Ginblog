package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	// AppMode 应用运行模式，如开发模式、生产模式等
	AppMode string
	// HttpPort HTTP服务监听的端口
	HttpPort string
	// JwtKey 用于JWT身份验证的密钥
	JwtKey string

	// 数据库配置
	// DbHost 数据库主机地址
	DbHost string
	// DbPort 数据库端口
	DbPort string
	// DbUser 数据库用户名
	DbUser string
	// DbPassWord 数据库密码
	DbPassWord string
	// DbName 数据库名称
	DbName string

	// Qiniu云存储配置
	// Zone 七牛云存储区域
	Zone int
	// AccessKey 七牛云存储的访问密钥
	AccessKey string
	// SecretKey 七牛云存储的秘密密钥
	SecretKey string
	// Bucket 七牛云存储的桶名称
	Bucket string
	// QiniuSever 七牛云存储的服务器地址
	QiniuSever string
)

// init 函数在包初始化时调用，用于加载配置文件
func init() {
	// 读取配置文件
	file, err := ini.Load("config/config.ini")
	if err != nil {
		// 如果读取配置文件出错，输出错误信息
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		return // 读取失败后停止执行
	}
	// 加载服务器相关配置
	LoadServer(file)
	// 加载数据库相关配置
	LoadData(file)
	// 加载七牛云存储相关配置
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	// 读取服务器模式 (AppMode)，默认为 "debug"
	AppMode = file.Section("server").Key("AppMode").MustString("debug")

	// 读取服务器端口 (HttpPort)，默认为 ":3000"
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")

	// 读取 JWT 密钥 (JwtKey)，默认为 "89js82js72"
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

// LoadData 从 ini 文件中加载数据库配置
func LoadData(file *ini.File) {
	// 读取数据库主机地址 (DbHost)，默认为 "localhost"
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")

	// 读取数据库端口 (DbPort)，默认为 "3306"
	DbPort = file.Section("database").Key("DbPort").MustString("3306")

	// 读取数据库用户名 (DbUser)，默认为 "ginblog"
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")

	// 读取数据库密码 (DbPassWord)，没有默认值
	DbPassWord = file.Section("database").Key("DbPassWord").String()

	// 读取数据库名称 (DbName)，默认为 "ginblog"
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

// LoadQiniu 从 ini 文件中加载七牛云配置
func LoadQiniu(file *ini.File) {
	// 读取七牛云的区域 (Zone)，默认为 1
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)

	// 读取七牛云的 AccessKey，没有默认值
	AccessKey = file.Section("qiniu").Key("AccessKey").String()

	// 读取七牛云的 SecretKey，没有默认值
	SecretKey = file.Section("qiniu").Key("SecretKey").String()

	// 读取七牛云的存储桶名称 (Bucket)，没有默认值
	Bucket = file.Section("qiniu").Key("Bucket").String()

	// 读取七牛云的服务器地址 (QiniuSever)，没有默认值
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
