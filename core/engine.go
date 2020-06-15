/*
基础查询引擎
提供初始化数据库连接、查询数量、执行SQL方法
*/

package core

import (
	"database/sql"
	"log"
)

// 查询引擎
type Engine struct {
	// 日志
	Logger Logger
	// 数据库连接
	DB *sql.DB
}

// 初始化数据库连接
func (engine *Engine) InitDB(driverName, dataSourceName string) *Engine {
	engine.Logger = &PrintLogger{}
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	engine.DB = db
	return engine
}

// 设置日志
func (engine *Engine) SetLogger(logger Logger) *Engine {
	engine.Logger = logger
	return engine
}

// 统计数量
func (engine *Engine) Count(sql string, params []interface{}) (int64, error) {
	// 执行查询
	engine.Logger.PrintInfo(sql, params)
	rows, err := engine.DB.Query(sql, params...)
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, err
	}
	var count int64 = 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			engine.Logger.PrintError(err)
			return -1, err
		}
	}
	return count, nil
}

// 执行SQL
func (engine *Engine) Exec(sql string, params []interface{}) (int64, int64, error) {
	// 执行SQL
	engine.Logger.PrintInfo(sql, params)
	result, err := engine.DB.Exec(sql, params...)
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	// 最后插入的ID
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	// 受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		engine.Logger.PrintError(err)
		return -1, -1, err
	}
	return lastInsertId, rowsAffected, nil
}
