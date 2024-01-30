package sqlmanage

import (
	"bytes"
	"database/sql"
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
// 	prefix TEXT NOT NULL,
// 	prefixEndPK TEXT NOT NULL
// )

var insertStmt *sql.Stmt

func Insert_Prefix(prefixName, prefix string) (err error) {
	// 生成预编译语句
	if insertStmt == nil {
		insertStmt, err = db.Prepare("INSERT INTO prefixCompare (prefixName, prefix, prefixEndPK) values (?, ?, ?)")
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
	res, err := insertStmt.Exec(prefixName, prefix, prefixEndPK.String())
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

func selectPrefixByPrefixName(prefixName string) (err error, prefix string, prefixEndPK string) {
	// 生成预编译语句
	if selectStmt == nil {
		selectStmt, err = db.Prepare("SELECT prefix, prefixEndPK FROM prefixCompare WHERE prefixName = ?")
		if err != nil {
			return
		}
	}

	rows, err := selectStmt.Query(prefixName)
	if err != nil {
		return
	}

	rows.Next()
	err = rows.Scan(&prefix, &prefixEndPK)
	return
}
