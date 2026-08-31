package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeromicro/go-zero/core/logx"
)

// DatabaseConf 数据库配置
type DatabaseConf struct {
	Host         string `json:",env=DATABASE_HOST"`                                                // 数据库地址
	Port         int    `json:",env=DATABASE_PORT"`                                                // 数据库端口
	Username     string `json:",default=root,env=DATABASE_USERNAME"`                               // 数据库用户名
	Password     string `json:",optional,env=DATABASE_PASSWORD"`                                   // 数据库密码
	DBName       string `json:",default=simple_admin,env=DATABASE_DBNAME"`                         // 数据库名称
	SSLMode      string `json:",optional,env=DATABASE_SSL_MODE"`                                   // PostgreSQL SSL 模式
	Type         string `json:",default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"` // 数据库类型
	MaxOpenConn  int    `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`                  // 最大打开连接数
	CacheTime    int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`                      // 缓存时间
	DBPath       string `json:",optional,env=DATABASE_DBPATH"`                                     // SQLite 数据库文件路径
	MysqlConfig  string `json:",optional,env=DATABASE_MYSQL_CONFIG"`                               // MySQL 额外连接参数
	PGConfig     string `json:",optional,env=DATABASE_PG_CONFIG"`                                  // PostgreSQL 额外连接参数
	SqliteConfig string `json:",optional,env=DATABASE_SQLITE_CONFIG"`                              // SQLite 额外连接参数
	Debug        bool   `json:",optional,env=DATABASE_DEBUG"`                                      // 是否开启数据库调试模式
}

// NewNoCacheDriver 创建不带缓存的 Ent 数据库驱动
func (c DatabaseConf) NewNoCacheDriver() *entsql.Driver {
	// 打开数据库连接
	db, err := sql.Open(c.Type, c.GetDSN())
	logx.Must(err)

	// 检查数据库连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	logx.Must(err)

	db.SetMaxOpenConns(c.MaxOpenConn)      // 最大打开连接数
	db.SetMaxIdleConns(20)                 // 最大空闲连接数
	db.SetConnMaxIdleTime(5 * time.Minute) // 连接最大空闲时间
	db.SetConnMaxLifetime(time.Hour)       // 连接最大生命周期

	// 创建 Ent 数据库驱动
	driver := entsql.OpenDB(c.Type, db)

	return driver
}

// MysqlDSN 获取 MySQL 连接地址
func (c DatabaseConf) MysqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=True%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.MysqlConfig,
	)
}

// PostgresDSN 获取 PostgreSQL 连接地址
func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.SSLMode,
		c.PGConfig,
	)
}

// SqliteDSN 获取 SQLite 连接地址
func (c DatabaseConf) SqliteDSN() string {
	if c.DBPath == "" {
		logx.Must(errors.New("SQLite 数据库文件路径不能为空"))
	}

	// 数据库文件不存在时自动创建
	if _, err := os.Stat(c.DBPath); os.IsNotExist(err) {
		f, err := os.OpenFile(c.DBPath, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			logx.Must(fmt.Errorf("创建 SQLite 数据库文件失败：%s", c.DBPath))
		}
		if err := f.Close(); err != nil {
			logx.Must(fmt.Errorf("关闭 SQLite 数据库文件失败：%s", c.DBPath))
		}
	} else {
		if err := os.Chmod(c.DBPath, 0660); err != nil {
			logx.Must(fmt.Errorf("unable to set permission code on %s: %v", c.DBPath, err))
		}
	}
	return fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1%s", c.DBPath, c.SqliteConfig)
}

// GetDSN 根据数据库类型获取对应连接地址
func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case "mysql":
		return c.MysqlDSN()
	case "postgres":
		return c.PostgresDSN()
	case "sqlite3":
		return c.SqliteDSN()
	default:
		return "mysql"
	}
}
