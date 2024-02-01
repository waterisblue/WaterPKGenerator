package pk

import (
	"pkgenerate/config"
	"strconv"
)

type PKManager struct {
	goroutineNum    int
	concurrentQueue *ConcurrentQueue
}

var isInit bool = false
var pkManager *PKManager

// 初始化PKManager
func InitManage() (pkManager *PKManager, err error) {
	if isInit {
		return
	}

	pkManager = &PKManager{
		goroutineNum:    0,
		concurrentQueue: NewConcurrentQueue(),
	}

	return
}

// 开始生成key
func Start() {
	configMap := config.GetConfig()
	minCache, _ := strconv.Atoi(configMap["pk.cache.minnum"])
	maxCache, _ := strconv.Atoi(configMap["pk.cache.num"])
	pkManager.goroutineNum = calculateGoroutineNum(maxCache)
	for {
		if pkManager.concurrentQueue.Len() < minCache {

		}
	}
}

func calculateGoroutineNum(maxCache int) int {
	if maxCache == 0 {
		return 0
	}
	if maxCache%36 != 0 {
		return maxCache/36 + 1
	}
}
