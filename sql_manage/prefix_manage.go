package sqlmanage

import (
	"bytes"
	"database/sql"
	"errors"
	"pkgenerate/config"
	"strconv"
)

var db *sql.DB

// 初始化数据库
func Init_db() (code int, err error) {
	db, err = sql.Open("sqllite3", "../sql/prefix.db")
	if err != nil {
		return 1, err
	}

	return 0, err
}

// 插入一个前缀，其中前缀通过prefixName对应
// 表结构：
// CREATE TABLE prefixCompare(
// 	id integer primary key autoincrement,
// 	prefixName TEXT NOT NULL UNIQUE,
// 	prefixEndPK TEXT NOT NULL
// )

var insertStmt *sql.Stmt

func Insert_Prefix(prefixName string) (err error) {
	// 查看该前缀是否已经存在
	isExist, _, _, _ := SelectPrefixByPrefixName(prefixName)
	if !isExist {
		// 该前缀已经存在，无法增加
		err = errors.New("该前缀" + prefixName + "已经存在")
		return
	}

	// 生成预编译语句
	if insertStmt == nil {
		insertStmt, err = db.Prepare("INSERT INTO prefixCompare (prefixName, prefixEndPK) values (?, ?, ?)")
		if err != nil {
			return
		}
	}
	var prefixEndPK bytes.Buffer

	// 前缀初始化
	configMap := config.GetConfig()
	pkLength, _ := strconv.Atoi(configMap["pk.length"])
	for i := 0; i < pkLength; i++ {
		prefixEndPK.WriteRune('0')
	}

	// 插入数据
	res, err := insertStmt.Exec(prefixName, prefixEndPK.String())
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()

	if err != nil || affect != 1 {
		return
	}

	// 正常执行
	return
}

// 根据前缀名查询前缀表内容
var selectStmt *sql.Stmt

func SelectPrefixByPrefixName(prefixName string) (isExist bool, prefixNameRes string, prefixEndPK string, err error) {
	// 生成预编译语句
	if selectStmt == nil {
		selectStmt, err = db.Prepare("SELECT prefixName, prefixEndPK FROM prefixCompare WHERE prefixName = ?")
		if err != nil {
			return
		}
	}

	rows, err := selectStmt.Query(prefixName)
	if err != nil {
		return
	}

	if !rows.Next() {
		isExist = false
		return
	}

	err = rows.Scan(&prefixNameRes, &prefixEndPK)
	return
}

// 查询前缀总数
var selectPrefixCount *sql.Stmt

func SelectPrefixCount() (count int, err error) {
	// 生成预编译语句
	if selectPrefixCount == nil {
		selectPrefixCount, err = db.Prepare("SELECT COUNT(1) as count FROM prefixCompare")
		if err != nil {
			return
		}
	}

	rows, err := selectPrefixCount.Query()
	if err != nil {
		return
	}

	rows.Next()
	err = rows.Scan(&count)
	return
}
