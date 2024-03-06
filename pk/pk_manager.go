package pk

import (
	"pkgenerate/config"
	sqlmanage "pkgenerate/sql_manage"
	"strconv"
)

type PKPrefixManager struct {
	pks      chan string
	isActive bool
}

var PKManager map[string]PKPrefixManager

// 开始生成key
func Start() {
	// 获取配置参数内容
	configMap := config.GetConfig()
	cacheNum, _ := strconv.Atoi(configMap["pk.cache.num"])
	// 查询前缀总数
	prefixCount, _ := sqlmanage.SelectPrefixCount()

	for i := 0; i < prefixCount; i++ {

	}
}
