package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var 创建下架表语句 = "CREATE TABLE IF NOT EXISTS xiajia (" +
	"hostname TEXT NOT NULL," +
	"ip TEXT NOT NULL," +
	"cpumax TEXT NOT NULL," +
	"memorymax TEXT NOT NULL," +
	"department TEXT NOT NULL," +
	"appname TEXT NOT NULL," +
	"count INTEGER NOT NULL," +
	"modified BOOLEAN DEFAULT 0," +
	"PRIMARY KEY (hostname, ip))"
var 创建降配表语句 = "CREATE TABLE IF NOT EXISTS jiangpei (" +
	"hostname TEXT NOT NULL," +
	"ip TEXT NOT NULL," +
	"cpumax TEXT NOT NULL," +
	"memorymax TEXT NOT NULL," +
	"department TEXT NOT NULL," +
	"appname TEXT NOT NULL," +
	"count INTEGER NOT NULL," +
	"modified BOOLEAN DEFAULT 0," +
	"PRIMARY KEY (hostname, ip))"
var 下架表变更位置false = "UPDATE xiajia SET modified = 0"
var 降配表变更位置false = "UPDATE jiangpei SET modified = 0"
var 下架表查询语句 = "SELECT * FROM xiajia WHERE hostname = ? AND ip = ?"
var 降配表查询语句 = "SELECT * FROM jiangpei WHERE hostname = ? AND ip = ?"
var 下架表插入一条新数据 = "INSERT INTO xiajia (hostname, ip, cpumax, memorymax, department, appname, count, modified) VALUES (?, ?, ?, ?, ?, ?, 1, 1)"
var 降配表插入一条新数据 = "INSERT INTO jiangpei (hostname, ip, cpumax, memorymax, department, appname, count, modified) VALUES (?, ?, ?, ?, ?, ?, 1, 1)"
var 下架表主机次数加1 = "UPDATE xiajia SET cpumax = ?, memorymax = ?, department = ?, appname = ?, count = ?, modified = 1 WHERE hostname = ? AND ip = ?"
var 降配表主机次数加1 = "UPDATE jiangpei SET cpumax = ?, memorymax = ?, department = ?, appname = ?, count = ?, modified = 1 WHERE hostname = ? AND ip = ?"
var 删除下架表中modified为0的语句 = "DELETE FROM xiajia WHERE modified = 0"
var 删除降配表中modified为0的语句 = "DELETE FROM jiangpei WHERE modified = 0"
var 删除下架表中上一期已下架或例外机器的语句 = "DELETE FROM xiajia WHERE hostname = ? AND ip = ?"
var 删除降配表中上一期已降配或例外机器的语句 = "DELETE FROM jiangpei WHERE hostname = ? AND ip = ?"
var 查询下架表所有信息的语句 = "SELECT * FROM xiajia"
var 查询降配表所有信息的语句 = "SELECT * FROM jiangpei"
var 查询表所有信息的语句 = "SELECT * FROM ? "

var host_db *sql.DB

func InitDB(path string) error {
	var err error
	host_db, err = sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	host_db.SetMaxOpenConns(1)
	return nil
}

func CloseDB() {
	host_db.Close()
}
