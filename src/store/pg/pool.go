package pg

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DataBase
//
//	@Description: 数据库链接
type DataBase struct {
	sql     *sql.DB
	gormSql *gorm.DB
}

// OpenDB
//
//	@Description: 创建数据库连接
//	@param dsn 数据库连接字符串
//	@param maxOpenConns 最大连接数
//	@param maxIdleConns 最大空闲连接数
//	@param connMaxLifetime 连接最大生命周期
//	@return *DataBase 数据库连接
//	@return error 错误信息
func OpenDB(dsn string, maxOpenConns, maxIdleConns int, connMaxLifetime time.Duration) (*DataBase, error) {
	// 创建数据库连接
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("连接 postgres 数据库失败: %w", err)
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	if e := sqlDB.Ping(); e != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("测试 postgres 连接失败: %w", err)
	}

	// 测试数据库时差
	var dbTime time.Time
	if e := sqlDB.QueryRow("SELECT NOW();").Scan(&dbTime); e != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("测试 postgres 时差失败: %w", err)
	}

	// 使用同连接创建 GORM 实例
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("初始化 GORM 失败：%w", err)
	}

	slog.Info("连接 postgres 数据库成功", "dsn_len", len(dsn), "time_difference", time.Now().Sub(dbTime))

	return &DataBase{
		sql:     sqlDB,
		gormSql: gormDB,
	}, nil
}
