package mysql

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Conf struct {
	// mysql 连接 dsn
	Dsn string
	// 最大开启的连接数
	MaxOpenConns int
	// 最大空闲的连接数
	MaxIdleConns int
	// 连接最长生命周期
	MaxLifeTime time.Duration
}

type MysqlClient interface {
	GetDB() *gorm.DB
}

type Client struct {
	db *gorm.DB
}

// 初始化 mysql 客户端
func NewClient(conf *Conf) (MysqlClient, error) {
	db, err := gorm.Open(mysql.Open(conf.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	rawdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	rawdb.SetMaxIdleConns(conf.MaxIdleConns)
	rawdb.SetMaxOpenConns(conf.MaxOpenConns)
	rawdb.SetConnMaxLifetime(conf.MaxLifeTime)
	return &Client{
		db: db,
	}, nil
}

func (c *Client) GetDB() *gorm.DB {
	return c.db
}
