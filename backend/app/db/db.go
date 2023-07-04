package db

import (
	"crypto/tls"
	"fmt"
	"time"

	gomysql "github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Env struct {
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
}

var (
	env Env
	db  *gorm.DB
)

type User struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	Email     string     `gorm:"type:VARCHAR(256);unique;not null"`
	Password  string     `gorm:"type:VARCHAR(128);not null"`
	CreatedAt *time.Time `gorm:"type:TIMESTAMP;autoCreateTime"`
	Token     string     `gorm:"type:TEXT"`
}

func getEnv(path string) Env {
	var env Env

	viper.SetConfigType("env")
	viper.SetConfigName("db")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Info("failed to read config", zap.Error(err))
		panic("failed to read config")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic("failed to unmarshal env")
	}

	return env
}

func InitDb(path string) {
	env = getEnv(path)

	gomysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: env.DbHost,
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/test?tls=tidb", env.DbUser, env.DbPassword, env.DbHost, env.DbPort)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
}

func Db() *gorm.DB {
	return db
}
