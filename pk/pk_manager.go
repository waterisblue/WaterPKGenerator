package pk

import (
	"fmt"
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
func Init() {
	// 获取配置参数内容
	configMap := config.GetConfig()
	cacheNum, _ := strconv.Atoi(configMap["pk.cache.num"])
	// 查询前缀总数
	prefixCount, _ := sqlmanage.SelectPrefixCount()

	PKManager = make(map[string]PKPrefixManager)

	for i := 1; i <= prefixCount; i++ {
		prefixName, prefixEndPK, err := sqlmanage.SelectPrefixById(i)
		if err != nil {
			fmt.Println("前缀名：" + prefixName + "生成错误，错误原因：" + err.Error())
			err = nil
			continue
		}

		// 初始化生成主键
		PKManager[prefixName] = PKPrefixManager{
			pks:      make(chan string, cacheNum),
			isActive: true,
		}
	}
}
