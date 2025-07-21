package config

import (
	"blog/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"Port"`
	} `yaml:"Server"`
	SecretKey string `yaml:"SecretKey"`
}

var (
	GormDB *gorm.DB
	SqlxDB *sqlx.DB
	Cfg    *Config
)

func LoadConfig(path string) (*Config, error) {
	var config Config

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func InitDb() error {
	log.Logger.Info("开始初始化数据库连接.....")
	config, err := LoadConfig("config/config.dev.yaml")
	Cfg = config
	if err != nil {
		log.Logger.Error("加载配置失败", zap.Error(err))
		return err
	}
	// 2. 使用配置中的 DSN 初始化数据库连接
	db, err := gorm.Open(mysql.Open(Cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Logger.Error("连接数据库失败", zap.Error(err))
		return fmt.Errorf("连接数据库失败: %v", err)
	}
	GormDB = db
	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Error("获取 sql.DB 失败", zap.Error(err))
		return fmt.Errorf("获取 sql.DB 失败: %v", err)
	}
	SqlxDB = sqlx.NewDb(sqlDB, "mysql")
	log.Logger.Info("✅ Successfully connected to database!")
	return nil
}
