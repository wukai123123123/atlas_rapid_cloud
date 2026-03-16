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

// DataBase 数据库连接
type DataBase struct {
	SQL  *sql.DB
	GORM *gorm.DB
}

// OpenDB 创建数据库连接
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

	slog.Warn("测试 postgres 时差", "sub_time", time.Now().Sub(dbTime))

	// 使用同连接创建 GORM 实例
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("初始化 GORM 失败：%w", err)
	}

	slog.Info("连接 postgres 数据库成功", "dsn_len", len(dsn))
	return &DataBase{
		SQL:  sqlDB,
		GORM: gormDB,
	}, nil
}
