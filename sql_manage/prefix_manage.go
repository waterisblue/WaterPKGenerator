package sqlmanage

import (
	"bytes"
	"database/sql"
	"errors"
	"pkgenerate/config"
	"pkgenerate/log"
	"strconv"
)

var db *sql.DB

func init() {
	prefixDB, err := sql.Open("sqllite3", "../sql/prefix.db")
	if err != nil {
		log.Error.Println(err.Error())
	}
	db = prefixDB
}

// 表结构：
// CREATE TABLE prefixCompare(
// 	id integer primary key autoincrement,
// 	prefixName TEXT NOT NULL UNIQUE,
// 	prefixEndPK TEXT NOT NULL
// )

var insertStmt *sql.Stmt

func InsertPrefix(prefixName string) (err error) {
	// 查看该前缀是否已经存在
	isExist, _, _, _ := SelectPrefixByPrefixName(prefixName)
	if !isExist {
		err = errors.New("该前缀" + prefixName + "已经存在")
		return
	}

	if insertStmt == nil {
		insertStmt, err = db.Prepare("INSERT INTO prefixCompare (prefixName, prefixEndPK) values (?, ?, ?)")
		if err != nil {
			return
		}
	}
	var prefixEndPK bytes.Buffer

	configMap := config.GetConfig()
	pkLength, _ := strconv.Atoi(configMap["pk.length"])
	for i := 0; i < pkLength; i++ {
		prefixEndPK.WriteRune('0')
	}

	res, err := insertStmt.Exec(prefixName, prefixEndPK.String())
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()

	if err != nil || affect != 1 {
		return
	}

	return
}

// 根据前缀名查询前缀表内容
var selectStmt *sql.Stmt

func SelectPrefixByPrefixName(prefixName string) (isExist bool, prefixNameRes string, prefixEndPK string, err error) {
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

// 根据id查询前缀表内容
// 查询id
var selectPrefixById *sql.Stmt

func SelectPrefixById(id int) (prefixName string, prefixEndPK string, err error) {
	if selectPrefixById == nil {
		selectPrefixById, err = db.Prepare("SELECT prefixName, prefixEndPK FROM prefixCompare WHERE id = ?")
		if err != nil {
			return
		}
	}

	rows, err := selectPrefixById.Query(id)
	if err != nil {
		return
	}

	rows.Next()
	err = rows.Scan(&prefixName, &prefixEndPK)
	return
}
