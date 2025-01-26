package dao

import (
	"context"
	"demo/model"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// 加载配置文件
func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
}

func InitMySQL() {
	var err error

	// 加载配置
	loadConfig()

	// 读取 MySQL 配置信息
	host := viper.GetString("MySQL.host")
	port := viper.GetString("MySQL.port")
	database := viper.GetString("MySQL.database")
	user := viper.GetString("MySQL.user")
	password := viper.GetString("MySQL.password")

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	// 连接 MySQL
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	log.Println("MySQL connected successfully!")

	// 自动迁移
	err = db.AutoMigrate(&model.User{}, &model.Content{})
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// 获取并初始化管理员用户
	username := viper.GetString("username")
	password = viper.GetString("password")
	phone := viper.GetString("phone")
	email := viper.GetString("email")

	// 添加管理员账户
	err = AddUser(username, password, email, phone, "admin")
	if err != nil {
		log.Fatalf("Error adding admin user: %v", err)
	}
}

func InitRedis() {
	// 加载配置
	loadConfig()

	// 读取 Redis 配置信息
	host := viper.GetString("Redis.host")
	port := viper.GetString("Redis.port")
	//password := viper.GetString("Redis.password")
	DB := viper.GetInt("Redis.DB")

	// 连接 Redis
	Rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: "",
		DB:       DB,
	})

	// 测试 Redis 连接
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected successfully!")
}
